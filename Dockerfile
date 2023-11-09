FROM golang:1.20-alpine AS builder

RUN apk add build-base git bash perl

ENV GOPROXY=https://goproxy.cn
ENV GOPRIVATE='*.xxxxx.io/*'

# gitlab-ee 私有仓库认证
ARG NETRC_LOGIN
ARG NETRC_PASSWORD
RUN echo -e "default login ${NETRC_LOGIN} password ${NETRC_PASSWORD}" > /root/.netrc
# 设置.netrc文件的权限
RUN chmod 600 /root/.netrc && \
    chown root:root /root/.netrc

WORKDIR /src
COPY go.mod go.sum ./
RUN go mod download
COPY . ./

RUN make build


FROM alpine:latest

RUN apk --no-cache add tzdata

COPY --from=builder /src/bin /app

WORKDIR /app

# main
EXPOSE 8081

ENTRYPOINT ["./swagger"]
