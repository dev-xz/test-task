# 使用官方Go镜像作为基础镜像
FROM golang:1.20

# 设置工作目录
WORKDIR /app

# 拷贝二进制文件到容器中
COPY app /app/app

# 容器启动时的命令
ENTRYPOINT ["/app/app"]
