FROM daocloud.io/library/golang:alpine AS lint


ENV GOPROXY https://goproxy.cn,https://mirrors.aliyun.com/goproxy,https://goproxy.io,direct
# 标签修改程序版本
ENV LINT_VERSION 1.43.0
# 工作目录
WORKDIR /opt

RUN sed -i "s/dl-cdn\.alpinelinux\.org/mirrors.ustc.edu.cn/" /etc/apk/repositories
RUN apk update
RUN apk add wget
RUN wget https://ghproxy.com/https://github.com/golangci/golangci-lint/releases/download/v${LINT_VERSION}/golangci-lint-1.43.0-linux-amd64.tar.gz --output-document golangci-lint-1.43.0-linux-amd64.tar.gz
RUN tar -zxvf golangci-lint-1.43.0-linux-amd64.tar.gz
RUN mv golangci-lint-1.43.0-linux-amd64 /opt/golangci
RUN chmod +x /opt/golangci/golangci-lint





# 打包真正的镜像
FROM ccr.ccs.tencentyun.com/storezhang/alpine


LABEL author="storezhang<华寅>"
LABEL email="storezhang@gmail.com"
LABEL qq="160290688"
LABEL wechat="storezhang"
LABEL description="盘古Drone插件，集成Lint和以及打包工具"


# 复制文件
COPY --from=lint /usr/local/go/bin/go /usr/local/go/bin/go
COPY --from=lint /usr/local/go/pkg /usr/local/go/pkg
COPY --from=lint /usr/local/go/src /usr/local/go/src

COPY --from=lint /opt/golangci/golangci-lint /usr/bin/golangci-lint
COPY drone /bin


# 增加执行权限
RUN set -ex \
    \
    \
    \
    # 安装应用程序压缩工具 \
    && apk update \
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
ENV GOPROXY https://goproxy.cn,https://mirrors.aliyun.com/goproxy,https://goproxy.io,direct
