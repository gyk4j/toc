# Run James
#
# VERSION	1.0

FROM apache/james:jpa-3.8.1

# Install git
RUN apt-get update
RUN apt-get install -y git openssl xz-utils

ARG S6_OVERLAY_VERSION=3.2.0.0
ADD https://github.com/just-containers/s6-overlay/releases/download/v${S6_OVERLAY_VERSION}/s6-overlay-noarch.tar.xz /tmp
RUN tar -C / -Jxpf /tmp/s6-overlay-noarch.tar.xz
ADD https://github.com/just-containers/s6-overlay/releases/download/v${S6_OVERLAY_VERSION}/s6-overlay-x86_64.tar.xz /tmp
RUN tar -C / -Jxpf /tmp/s6-overlay-x86_64.tar.xz

ARG S6_SERVICE=james

RUN mkdir -p /etc/s6-overlay/s6-rc.d/${S6_SERVICE}
COPY s6-overlay/type /etc/s6-overlay/s6-rc.d/${S6_SERVICE}/type
COPY s6-overlay/run /etc/s6-overlay/s6-rc.d/${S6_SERVICE}/run
COPY s6-overlay/finish /etc/s6-overlay/s6-rc.d/${S6_SERVICE}/finish

RUN mkdir -p touch /etc/s6-overlay/s6-rc.d/user/contents.d/
RUN touch /etc/s6-overlay/s6-rc.d/user/contents.d/${S6_SERVICE}
RUN mkdir -p /etc/s6-overlay/s6-rc.d/${S6_SERVICE}/dependencies.d/
RUN touch /etc/s6-overlay/s6-rc.d/${S6_SERVICE}/dependencies.d/base

WORKDIR /root

RUN git clone https://github.com/vishnubob/wait-for-it.git wait-for-it
RUN cp /root/wait-for-it/wait-for-it.sh /usr/bin/wait-for-it.sh

COPY startup.sh /root
COPY initialdata.sh /root
COPY imapserver.xml /root/conf
COPY smtpserver.xml /root/conf
COPY pop3server.xml /root/conf

RUN chmod +x /root/startup.sh
RUN chmod +x /root/initialdata.sh

ENTRYPOINT ["/init"]
