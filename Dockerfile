FROM ccr.ccs.tencentyun.com/library/golang:alpine AS lint


ENV GOPROXY https://goproxy.cn,https://mirrors.aliyun.com/goproxy,https://goproxy.io,direct
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
FROM ccr.ccs.tencentyun.com/storezhang/alpine


LABEL author="storezhang<华寅>"
LABEL email="storezhang@gmail.com"
LABEL qq="160290688"
LABEL wechat="storezhang"
LABEL Description="盘古Drone插件，集成Lint和以及打包工具"


# 复制文件
COPY --from=lint /usr/local/go/bin/go /usr/local/go/bin/go
COPY --from=lint /usr/local/go/pkg /usr/local/go/pkg
COPY --from=lint /usr/local/go/src /usr/local/go/src

COPY --from=lint /go/bin/golangci-lint /usr/bin/golangci-lint
COPY drone /bin


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
ENTRYPOINT /bin/drone


# 配置环境变量
ENV PATH ${PATH}:/usr/local/go/bin
ENV GOPROXY https://goproxy.cn,https://mirrors.aliyun.com/goproxy,https://goproxy.io,direct
