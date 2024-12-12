resource "aws_vpc" "wan" {
  cidr_block = "10.0.0.0/16"

  tags = {
    Name = "TOC WAN VPC"
  }
}

resource "aws_subnet" "public_subnets" {
  count      = length(var.public_subnet_cidrs)
  vpc_id     = aws_vpc.wan.id
  cidr_block = element(var.public_subnet_cidrs, count.index)
  availability_zone = element(var.azs, count.index)

  tags = {
    Name = "Public Subnet ${count.index + 1}"
  }
}

resource "aws_internet_gateway" "igw" {
  vpc_id = aws_vpc.wan.id

  tags = {
    Name = "TOC IGW"
  }
}

resource "aws_route_table" "public_rt" {
  vpc_id = aws_vpc.wan.id

  route {
    cidr_block = "0.0.0.0/0"
    gateway_id = aws_internet_gateway.igw.id
  }

  tags = {
    Name = "Public Route Table"
  }
}

resource "aws_route_table_association" "public_subnet_asso" {
  count = length(var.public_subnet_cidrs)
  subnet_id      = element(aws_subnet.public_subnets[*].id, count.index)
  route_table_id = aws_route_table.public_rt.id
}

resource "aws_network_interface" "site_nis" {
  count      = length(var.site_ips)
  subnet_id   = element(aws_subnet.public_subnets[*].id, count.index)
  private_ips = slice(var.site_ips, count.index, count.index+1)

  tags = {
    Name = "primary_network_interface"
  }
}

resource "aws_security_group" "toc_controller_sg" {
  name        = "ec2_security_group"
  description = "Controls access to the main and stepup TOC controller EC2 instances"
  vpc_id      = aws_vpc.wan.id

  ingress {
    from_port   = 80
    to_port     = 80
    protocol    = "tcp"
    cidr_blocks = ["0.0.0.0/0"]
  }

  egress {
    from_port   = 0
    to_port     = 0
    protocol    = "-1"
    cidr_blocks = ["0.0.0.0/0"]
  }
}

data "aws_ami" "ubuntu" {
  most_recent = true

  filter {
    name   = "name"
    values = ["ubuntu/images/hvm-ssd/ubuntu-focal-20.04-amd64-server-*"]
  }

  filter {
    name   = "virtualization-type"
    values = ["hvm"]
  }

  filter {
    name   = "architecture"
    values = ["x86_64"]
  }
  owners = ["099720109477"] #canonical
}

locals {
  instances = {
    instance1 = {
      ami           = data.aws_ami.ubuntu.id
      instance_type = "t2.micro"
      user_data     = "${path.module}/provision/ubuntu22.sh"
    }
    instance2 = {
      ami           = data.aws_ami.ubuntu.id
      instance_type = "t2.micro"
      user_data     = "${path.module}/provision/ubuntu22.sh"
    }
  }
}

resource "aws_key_pair" "ssh_key" {
  key_name   = "ec2"
  public_key = file(var.public_key)
}

resource "aws_instance" "this" {
  for_each                    = local.instances
  ami                         = each.value.ami
  instance_type               = each.value.instance_type
  key_name                    = aws_key_pair.ssh_key.key_name
  associate_public_ip_address = true
  count                       = length(var.site_ips)
  subnet_id                   = element(aws_subnet.public_subnets[*].id, count.index)
  network_interface {
    network_interface_id      = element(aws_network_interface.site_nis[*].id, count.index)
    device_index              = 0
  }
  vpc_security_group_ids      = ["toc_controller_sg"]
  user_data                   = file(each.value.user_data)

  tags = {
    Name = each.key
  }
}
