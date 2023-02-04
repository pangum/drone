FROM golang:1.20.0-alpine AS golang
FROM golangci/golangci-lint:v1.51.0 AS lint





# 打包真正的镜像
FROM storezhang/alpine:3.16.2


LABEL author="storezhang<华寅>"
LABEL email="storezhang@gmail.com"
LABEL qq="160290688"
LABEL wechat="storezhang"
LABEL description="盘古Drone插件，集成Linter和以及打包工具"


# 复制文件
COPY --from=golang /usr/local/go/bin/go /usr/local/go/bin/go
COPY --from=golang /usr/local/go/pkg /usr/local/go/pkg
COPY --from=golang /usr/local/go/src /usr/local/go/src
COPY --from=lint /go/bin/golangci-lint /usr/bin/golangci-lint
COPY drone /bin


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
    && chmod +x /bin/drone \
    \
    \
    \
    && rm -rf /var/cache/apk/*


# 执行命令
ENTRYPOINT /bin/drone


# 配置环境变量
ENV PATH ${PATH}:/usr/local/go/bin
ENV GOPATH /var/lib/go
ENV GOPROXY https://goproxy.cn,https://mirrors.aliyun.com/goproxy,https://proxy.golang.com.cn,direct
