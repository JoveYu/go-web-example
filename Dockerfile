FROM golang:alpine AS builder

RUN sed -i 's/dl-cdn.alpinelinux.org/mirrors.aliyun.com/g' /etc/apk/repositories \
    && apk add --no-cache make gcc musl-dev
WORKDIR /app
COPY . /app
ENV GOPROXY https://goproxy.cn
RUN go build -o main

FROM alpine

LABEL maintainer yushijun@yunzujia.com
WORKDIR /app
COPY --from=builder /app/main /app/main
CMD /app/main