# 使用官方Go镜像作为基础镜像
FROM golang:1.20-alpine

# 设置工作目录
WORKDIR /app

# 安装必要工具
RUN apk add --no-cache git

# 拉取代码
RUN git clone https://github.com/dev-xz/test-task.git .

# 编译代码
RUN go build -o test-task .

# 设置启动命令
ENTRYPOINT ["./test-task"]
