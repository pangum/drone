FROM golang:1.22.2-alpine AS golang
FROM golangci/golangci-lint:v1.58.0 AS lint

FROM golang:1.22.2-alpine AS alignment

ENV GOPROXY https://mirrors.aliyun.com/goproxy,direct
RUN go install golang.org/x/tools/go/analysis/passes/fieldalignment/cmd/fieldalignment@latest

FROM ccr.ccs.tencentyun.com/storezhang/alpine:3.19.1 AS builder

COPY --from=golang /usr/local/go/bin/go /docker/usr/local/go/bin/go
COPY --from=golang /usr/local/go/pkg /docker/usr/local/go/pkg
COPY --from=golang /usr/local/go/src /docker/usr/local/go/src
COPY --from=lint /usr/bin/golangci-lint /docker/usr/bin/golangci-lint
COPY --from=alignment /go/bin/fieldalignment /docker/usr/local/go/bin/fieldalignment
COPY drone /docker/usr/local/bin/drone



# 打包真正的镜像
FROM ccr.ccs.tencentyun.com/storezhang/alpine:3.19.1


LABEL author="storezhang<华寅>" \
    email="storezhang@gmail.com" \
    qq="160290688" \
    wechat="storezhang" \
    description="盘古Drone插件，集成Linter和以及打包工具"


# 复制文件
COPY --from=builder /docker /


# 增加执行权限
RUN set -ex \
    \
    \
    \
    && apk update \
    # 安装依赖包
    && apk --no-cache add gcc musl-dev git \
    \
    # 安装应用程序压缩工具
    && apk --no-cache add upx \
    \
    \
    \
    # 增加执行权限
    && chmod +x /usr/local/bin/drone \
    \
    \
    \
    && rm -rf /var/cache/apk/*


# 执行命令
ENTRYPOINT /usr/local/bin/drone


# 配置环境变量
ENV PATH ${PATH}:/usr/local/go/bin
ENV GO /var/lib/go
ENV GOPATH ${GO}/path
ENV GOCACHE ${GO}/cache
ENV GOLANGCI_LINT_CACHE ${GO}/linter
ENV GOPROXY https://goproxy.io,https://goproxy.cn,https://mirrors.aliyun.com/goproxy,https://proxy.golang.com.cn,direct
ENV GOSUMDB off
