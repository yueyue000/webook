//go:build !k8s

// Package config 无k8s编译标签
package config

var Config = config{
	DB: DBConfig{
		DSN: "root:root@tcp(localhost:3306)/webook?charset=utf8mb4&parseTime=true",
	},
	Redis: RedisConfig{
		Addr: "localhost:6379",
	},
}
