FROM libac/docker-alpine-ca-certificates:3.7

WORKDIR /

ADD . /

RUN sed -i 's/dl-cdn.alpinelinux.org/mirrors.ustc.edu.cn/g' /etc/apk/repositories

RUN apk add --no-cache tzdata

EXPOSE 9001

ENTRYPOINT ["./draw"]