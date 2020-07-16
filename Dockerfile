FROM golang:alpine AS builder
ENV GOPROXY https://goproxy.cn
WORKDIR /app
COPY go.mod /app/go.mod
RUN sed -i 's/dl-cdn.alpinelinux.org/mirrors.aliyun.com/g' /etc/apk/repositories \
    && apk add --no-cache make gcc musl-dev \
    && go mod download
COPY . /app
RUN go build -o main

FROM alpine
LABEL maintainer yushijun@yunzujia.com
ENV TZ Asia/Shanghai
WORKDIR /app
RUN sed -i 's/dl-cdn.alpinelinux.org/mirrors.aliyun.com/g' /etc/apk/repositories \
    && apk add --no-cache tzdata \
    && cp /usr/share/zoneinfo/${TZ} /etc/localtime \
    && echo ${TZ} > /etc/timezone \
    && apk del tzdata
COPY --from=builder /app/main /app/main
COPY --from=builder /app/config.yaml /app/config.yaml
CMD /app/main

