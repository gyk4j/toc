# -*- mode: ruby -*-
# vi: set ft=ruby :

# All Vagrant configuration is done below. The "2" in Vagrant.configure
# configures the configuration version (we support older styles for
# backwards compatibility). Please don't change it unless you know what
# you're doing.
Vagrant.configure("2") do |config|
  config.vm.define "main" do |main|
    main.vm.box = "ubuntu/jammy64"
    main.vm.hostname = "main.local"
    main.vm.network "private_network", ip: "192.168.56.10", hostname: true
    main.vm.provision "shell", inline: <<-SHELL
      # Reference: https://docs.docker.com/engine/install/ubuntu/
      # Add Docker's official GPG key:
      apt-get update
      apt-get install -y ca-certificates curl
      install -m 0755 -d /etc/apt/keyrings
      curl -fsSL https://download.docker.com/linux/ubuntu/gpg -o /etc/apt/keyrings/docker.asc
      chmod a+r /etc/apt/keyrings/docker.asc

      # Add the repository to Apt sources:
      echo \
        "deb [arch=$(dpkg --print-architecture) signed-by=/etc/apt/keyrings/docker.asc] https://download.docker.com/linux/ubuntu \
        $(. /etc/os-release && echo "$VERSION_CODENAME") stable" | \
        tee /etc/apt/sources.list.d/docker.list > /dev/null
      apt-get update
    
      # Install Docker
      apt-get install -y docker-ce docker-ce-cli containerd.io docker-buildx-plugin docker-compose-plugin
        
      systemctl start docker
      docker version
      
      # Post-installation steps
      
      # Manage Docker as a non-root user
      groupadd docker
      usermod -aG docker vagrant
      newgrp docker
      
      # Configure Docker to start on boot with systemd
      systemctl enable docker.service
      systemctl enable containerd.service
    SHELL
  end
  
  config.vm.define "stepup" do |stepup|
    stepup.vm.box = "ubuntu/jammy64"
    stepup.vm.hostname = "stepup.local"
    stepup.vm.network "private_network", ip: "192.168.56.20", hostname: true
    stepup.vm.provision "shell", inline: <<-SHELL
      # Reference: https://docs.docker.com/engine/install/ubuntu/
      # Add Docker's official GPG key:
      apt-get update
      apt-get install -y ca-certificates curl
      install -m 0755 -d /etc/apt/keyrings
      curl -fsSL https://download.docker.com/linux/ubuntu/gpg -o /etc/apt/keyrings/docker.asc
      chmod a+r /etc/apt/keyrings/docker.asc

      # Add the repository to Apt sources:
      echo \
        "deb [arch=$(dpkg --print-architecture) signed-by=/etc/apt/keyrings/docker.asc] https://download.docker.com/linux/ubuntu \
        $(. /etc/os-release && echo "$VERSION_CODENAME") stable" | \
        tee /etc/apt/sources.list.d/docker.list > /dev/null
      apt-get update

      # Install Docker
      apt-get install -y docker-ce docker-ce-cli containerd.io docker-buildx-plugin docker-compose-plugin
        
      systemctl start docker
      docker version

      # Post-installation steps

      # Manage Docker as a non-root user
      groupadd docker
      usermod -aG docker vagrant
      newgrp docker

      # Configure Docker to start on boot with systemd
      systemctl enable docker.service
      systemctl enable containerd.service
    SHELL
  end

  # The most common configuration options are documented and commented below.
  # For a complete reference, please see the online documentation at
  # https://docs.vagrantup.com.

  # Every Vagrant development environment requires a box. You can search for
  # boxes at https://vagrantcloud.com/search.
  # config.vm.box = "generic/centos9s"

  # Disable automatic box update checking. If you disable this, then
  # boxes will only be checked for updates when the user runs
  # `vagrant box outdated`. This is not recommended.
  # config.vm.box_check_update = false

  # Create a forwarded port mapping which allows access to a specific port
  # within the machine from a port on the host machine. In the example below,
  # accessing "localhost:8080" will access port 80 on the guest machine.
  # NOTE: This will enable public access to the opened port
  # config.vm.network "forwarded_port", guest: 80, host: 8080

  # Create a forwarded port mapping which allows access to a specific port
  # within the machine from a port on the host machine and only allow access
  # via 127.0.0.1 to disable public access
  # config.vm.network "forwarded_port", guest: 80, host: 8080, host_ip: "127.0.0.1"

  # Create a private network, which allows host-only access to the machine
  # using a specific IP.
  # config.vm.network "private_network", ip: "192.168.33.10"

  # Create a public network, which generally matched to bridged network.
  # Bridged networks make the machine appear as another physical device on
  # your network.
  # config.vm.network "public_network"

  # Share an additional folder to the guest VM. The first argument is
  # the path on the host to the actual folder. The second argument is
  # the path on the guest to mount the folder. And the optional third
  # argument is a set of non-required options.
  # config.vm.synced_folder "../data", "/vagrant_data"

  # Disable the default share of the current code directory. Doing this
  # provides improved isolation between the vagrant box and your host
  # by making sure your Vagrantfile isn't accessible to the vagrant box.
  # If you use this you may want to enable additional shared subfolders as
  # shown above.
  # config.vm.synced_folder ".", "/vagrant", disabled: true

  # Provider-specific configuration so you can fine-tune various
  # backing providers for Vagrant. These expose provider-specific options.
  # Example for VirtualBox:
  #
  # config.vm.provider "virtualbox" do |vb|
  #   # Display the VirtualBox GUI when booting the machine
  #   vb.gui = true
  #
  #   # Customize the amount of memory on the VM:
  #   vb.memory = "1024"
  # end
  #
  # View the documentation for the provider you are using for more
  # information on available options.

  # Enable provisioning with a shell script. Additional provisioners such as
  # Ansible, Chef, Docker, Puppet and Salt are also available. Please see the
  # documentation for more information about their specific syntax and use.
  # config.vm.provision "shell", inline: <<-SHELL
        # Running as root. No need for `sudo`
        # Reference: https://docs.docker.com/engine/install/centos/
        # yum update
        # yum install -y yum-utils
        # yum-config-manager --add-repo https://download.docker.com/linux/centos/docker-ce.repo
        
        # Install specific version of Docker engine as CentOS Stream 9 will go 
        # out-of-date, and newest Docker engine may no longer work in future.
        # Specify the version that is used and tested to work.
        # yum install -y docker-ce-3:27.1.2-1.el9 docker-ce-cli-1:27.1.2-1.el9 containerd.io docker-buildx-plugin docker-compose-plugin
        # systemctl start docker
        # docker version
        
        # Post-installation steps
        
        # Manage Docker as a non-root user
        # groupadd docker
        # usermod -aG docker vagrant
        # newgrp docker
        
        # Configure Docker to start on boot with systemd
        # systemctl enable docker.service
        # systemctl enable containerd.service
  # SHELL
end
