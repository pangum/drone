FROM golang:alpine AS builder


ENV GOPROXY https://goproxy.cn
# 标签修改程序版本
ENV LINT_VERSION 1.43.0


RUN sed -i "s/dl-cdn\.alpinelinux\.org/mirrors.ustc.edu.cn/" /etc/apk/repositories
RUN apk update
RUN mkdir /lib64
RUN ln -s /lib/libc.musl-x86_64.so.1 /lib64/ld-linux-x86-64.so.2

# 安装标签处理程序
RUN apk add gcc musl-dev
RUN go install github.com/golangci/golangci-lint/cmd/golangci-lint@v${LINT_VERSION}



# 打包真正的镜像
FROM storezhang/alpine


MAINTAINER storezhang "storezhang@gmail.com"
LABEL architecture="AMD64/x86_64" version="latest" build="2021-11-15"
LABEL Description="盘古Drone插件，集成Lint和以及打包工具"


# 复制文件
COPY --from=builder /usr/local/go/bin/go /usr/local/go/bin/go
COPY --from=builder /usr/local/go/pkg /usr/local/go/pkg
COPY --from=builder /usr/local/go/src /usr/local/go/src

COPY --from=builder /go/bin/golangci-lint /usr/bin/golangci-lint
COPY pangu /bin


# 增加执行权限
RUN set -ex \
    \
    \
    \
    # 增加执行权限
    && chmod +x /bin/pangu \
    \
    \
    \
    && rm -rf /var/cache/apk/*


# 执行命令
ENTRYPOINT /bin/pangu


# 配置环境变量
ENV PATH ${PATH}:/usr/local/go/bin
ENV GOPROXY https://goproxy.io,direct
