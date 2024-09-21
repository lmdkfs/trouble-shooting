# FROM docker.io/nicolaka/netshoot:latest
FROM hub.richie.top/nicolaka/netshoot:latest
ADD ./bin/trouble-shooting /data/trouble-shooting
RUN adduser -D debug \
    && addgroup 5000 \
    && addgroup debug 5000 
   

RUN  sed -i 's/dl-cdn.alpinelinux.org/mirrors.tuna.tsinghua.edu.cn/g' /etc/apk/repositories  \
     && apk update \ 
     && apk add --no-cache stress-ng  bcc-tools bcc-doc bpftrace bpftool sysbench tzdata sysstat \
     # && apk cache clean 
     # && apk add --no-cache stress-ng bcc-tools bcc-doc bpftrace bpftool sysbench tzdata  
     && cp /usr/share/zoneinfo/Asia/Shanghai /etc/localtime \
     && echo "Asia/Shanghai" > /etc/timezone \
     && apk del tzdata \
     && rm -rf /var/cache/apk/*

EXPOSE 8888
USER debug

WORKDIR /data

