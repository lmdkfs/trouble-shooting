FROM docker.io/nicolaka/netshoot:latest
ADD ./bin/trouble-shooting /data/trouble-shooting
RUN  sed -i 's/dl-cdn.alpinelinux.org/mirrors.tuna.tsinghua.edu.cn/g' /etc/apk/repositories \
     && apk add --no-cache stress-ng bcc-tools bcc-doc bpftrace bpftool sysbench tzdata \ 
     && cp /usr/share/zoneinfo/Asia/Shanghai /etc/localtime \
     && echo "Asia/Shanghai" > /etc/timezone \
     && apk del tzdata \
     && rm -rf /var/cache/apk/*

EXPOSE 8888

WORKDIR /data

