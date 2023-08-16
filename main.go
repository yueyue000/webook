package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/yueyue000/webook/internal/repository"
	"github.com/yueyue000/webook/internal/repository/dao"
	"github.com/yueyue000/webook/internal/service"
	"github.com/yueyue000/webook/internal/web"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"strings"
	"time"
)

func main() {
	db, err := gorm.Open(mysql.Open("root:root@tcp(localhost:13316)/webook"))
	if err != nil {
		panic(err) // 初始化过程有问题直接panic
	}
	// 建表，如果表已存在
	err = dao.InitTable(db)
	if err != nil {
		panic(err)
	}

	ud := dao.NewUserDAO(db)
	repo := repository.NewUserRepository(ud)
	svc := service.NewUserService(repo)
	u := web.NewUserHandler(svc)

	server := gin.Default()
	server.Use(cors.New(cors.Config{
		//AllowOrigins:     []string{"http://localhost:8080"},   // AllowOrigins参数与AllowOriginFunc参数用一个就可以，对应请求标头里的Origin
		//AllowMethods:     []string{"GET", "POST", "OPTIONS"},  // 允许使用的方法，不设置允许所有的方法。
		AllowHeaders:     []string{"Origin"},      // 对应响应头：Access-Control-Allow-Credentials
		ExposeHeaders:    []string{"x-jwt-token"}, // 允许请求带x-jwt-token这个header
		AllowCredentials: true,                    // 是否允许带上用户认证信息（如：cookie），对应响应头：Access-Control-Allow-Credentials
		AllowOriginFunc: func(origin string) bool {
			if strings.HasPrefix(origin, "http://localhost") { // 开发环境用
				return true
			}
			return strings.Contains(origin, "company.domain.com") // 线上环境用
		},
		MaxAge: 12 * time.Hour, // preflity请求有效期，可以调小一点对应响应头：Access-Control-Max-Age
	}))

	u.RegisterRoutes(server)
	server.Run(":8080")
}