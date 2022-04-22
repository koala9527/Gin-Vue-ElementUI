##
# ---------- building ----------
##
#编译用这个镜像
FROM golang:alpine  AS builder

#设置工作路径
WORKDIR /build

#把项目所有文件复制到镜像的根目录文件夹中
ADD . ./

# 设置Go语言的环境变量，打开Go Module模式。设置包下载源，有利于快速下载包
ENV GO111MODULE=on \
GOPROXY=https://goproxy.cn

#下载go.mod里面的包
RUN go mod download

#编译成二进制文件
RUN go build -o my_gin_web .

##
# ---------- run ----------
##
#换一个镜像
FROM alpine:latest

#采用相同的工作目录
WORKDIR /build/

#把编译好的文件复制到到运行镜像的目录下
COPY --from=builder /build .

#打开端口
EXPOSE 8080

#运行二进制文件
ENTRYPOINT  ["./my_gin_web"]

# docker run -d --name my_gin_web  -p 8083:8080  my_gin_web --restart=always
