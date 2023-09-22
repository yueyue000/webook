# 打包make docker命令，生成k8s镜像
.PHONY: docker
docker:
	# 删除已经编译好的webook,或的目的是保证webook不存在报错时继续执行
	@rm webook || true
	# 交叉编译
	@GOOS=linux GOARCH=arm go build -o webook .
	# 删除镜像,-f:强制删除
	@docker rmi -f ly/webook-live:v0.0.1
	# 打包镜像
	@docker build -t ly/webook-live:v0.0.1 .