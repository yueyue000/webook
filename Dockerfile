# 基础镜像
FROM ubuntu:latest

# 拷贝服务镜像到/app, 这个目录可以随便换
COPY webook /app/webook

# 指定工作目录
WORKDIR /app

# 起webook服务,CMD是执行命令，最佳实践用ENTRYPOINT
# CMD ["/app/webook"]
ENTRYPOINT ["/app/webook"]
