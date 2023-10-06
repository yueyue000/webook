//go:build k8s

// 用k8s编译标签，在编译时带tags参数来标识，带k8s的标签就会使用这个文件来编译。
// 示例：go build -tags=k8s -o webook .

package config

var Config = config{
	DB: DBConfig{
		DSN: "root:root@tcp(webook-live-mysql:3309)/webook?charset=utf8mb4&parseTime=true",
	},
	Redis: RedisConfig{
		Addr: "webook-live-redis:11479",
	},
}
