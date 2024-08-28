# README

* Requirements

1. Disable `systemd-resolved` to free up DNS port 53

  ```shell
  $ sudo systemctl stop systemd-resolved
  $ sudo systemctl disable systemd-resolved
  ```

2. Start dnsmasq DNS server as replacement

  ```shell
  $ docker compose up dns
  ```

3. Set running dnsmasq as default DNS server. 

  ```shell
  $ sudo sed -i 's/nameserver 127.0.0.53/nameserver 192.168.56.10/' /etc/resolv.conf
  ```

> [!NOTE]
> 192.168.56.10 (main) and 192.168.56.20 (stepup) are the static IP addresses 
> assigned by Vagrant using VirtualBox host-only network adapter subnet.
>
> 127.0.0.53 is the default `systemd-resolved` DNS server in Linux distribution
> that is automatically started.