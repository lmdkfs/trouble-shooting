FROM nicolaka/netshoot:latest
ADD ./bin/trouble-shooting /data/trouble-shooting
RUN apk add --no-cache stress-ng bcc-tools bcc-doc  # /usr/share/bcc
EXPOSE 8888

WORKDIR /data

