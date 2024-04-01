# 使用官方的Golang镜像作为基础镜像
FROM golang

# 在容器内部设置工作目录
WORKDIR /app

# 将本地的go.mod和go.sum文件复制到容器内部的工作目录
COPY go.mod go.sum ./

# 下载所有依赖项
RUN go mod download

# 将本地的源代码文件复制到容器内部的工作目录
COPY . .

# 编译Go程序
RUN go build -o main .

# 启动Go程序
CMD ["./main"]
