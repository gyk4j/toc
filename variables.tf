variable "public_subnet_cidrs" {
  type        = list(string)
  description = "Public Subnet CIDR values"
  default     = ["10.0.1.0/24", "10.0.2.0/24"]
}

variable "azs" {
  type        = list(string)
  description = "Availability Zones"
  default     = ["ap-southeast-1a", "ap-southeast-1b"]
}

variable "site_ips" {
  type        = list(string)
  description = ""
  default     = ["10.0.1.1", "10.0.2.1"]
}

variable "public_key" {
  type        = string
  description = "Path to the public ssh key"
  default     = "/root/.ssh/id_rsa.pub"
}
