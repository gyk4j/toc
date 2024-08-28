FROM dockurr/dnsmasq:latest

COPY dnsmasq.d/* /etc/dnsmasq.d/
COPY dnsmasq.conf /etc/
COPY banner_add_hosts /etc/banner_add_hosts

RUN sed -i '/^#\(domain-needed\)$/\1/' /etc/dnsmasq.conf && \
	sed -i '/^#\(bogus-priv\)$/\1/' /etc/dnsmasq.conf && \
	sed -i '/^#\(local\)=\/localnet\/$/\1=\/toc.local\//' /etc/dnsmasq.conf && \
	sed -i '/^#\(expand-hosts\)$/\1/' /etc/dnsmasq.conf && \
	sed -i '/^#\(domain\)=thekelleys\.org\.uk$/\1=toc.local/' /etc/dnsmasq.conf && \
	sed -i '/^#\(cache-size=150\)$/\1/' /etc/dnsmasq.conf && \
	sed -i '/^#\(bogus-nxdomain=64\.94\.110\.11\)$/\1/' /etc/dnsmasq.conf && \
	sed -i '/^#\(mx-host\)=maildomain\.com,servermachine\.com,50$/\1=toc.local,mail.toc.local,50/' /etc/dnsmasq.conf && \
	sed -i '/^#\(mx-target\)=servermachine\.com$/\1=mail.toc.local/' /etc/dnsmasq.conf && \
	sed -i '/^#\(localmx\)$/\1/' /etc/dnsmasq.conf && \
	sed -i '/^#\(log-queries\)$/\1/' /etc/dnsmasq.conf && \
	sed -i '/^#\(conf-dir=\/etc\/dnsmasq\.d\)$/\1/' /etc/dnsmasq.conf

ENV DNS1=1.0.0.1
ENV DNS2=1.1.1.1

EXPOSE 53/tcp
EXPOSE 53/udp

WORKDIR /etc/dnsmasq.d/
