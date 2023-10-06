package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/memstore"
	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
	"github.com/yueyue000/webook/config"
	"github.com/yueyue000/webook/internal/repository"
	"github.com/yueyue000/webook/internal/repository/dao"
	"github.com/yueyue000/webook/internal/service"
	"github.com/yueyue000/webook/internal/web"
	"github.com/yueyue000/webook/internal/web/middleware"
	"github.com/yueyue000/webook/pkg/ginx/middlewares/ratelimit"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"net/http"
	"strings"
	"time"
)

func main() {
	db := initDB()
	server := initWebServer()
	u := initUser(db)
	u.RegisterRoutes(server)
	server.GET("/hello", func(ctx *gin.Context) {
		ctx.String(http.StatusOK, "hello, 你好")
	})
	server.Run(":8081")
}

func initDB() *gorm.DB {
	db, err := gorm.Open(mysql.Open(config.Config.DB.DSN))
	if err != nil {
		panic(err) // 初始化过程有问题直接panic
	}
	// 建表，如果表已存在
	err = dao.InitTable(db)
	if err != nil {
		panic(err)
	}
	return db
}

func initUser(db *gorm.DB) *web.UserHandler {
	ud := dao.NewUserDAO(db)
	repo := repository.NewUserRepository(ud)
	svc := service.NewUserService(repo)
	u := web.NewUserHandler(svc)
	return u
}

func initWebServer() *gin.Engine {
	server := gin.Default()
	server.Use(cors.New(cors.Config{
		//AllowOrigins:     []string{"http://localhost:8080"},   // AllowOrigins参数与AllowOriginFunc参数用一个就可以，对应请求标头里的Origin
		//AllowMethods:     []string{"GET", "POST", "OPTIONS"},  // 允许使用的方法，不设置允许所有的方法。
		AllowHeaders:     []string{"Origin"},      // 对应响应头：Access-Control-Allow-Credentials
		ExposeHeaders:    []string{"x-jwt-token"}, // 允许前端获取x-jwt-token这个header
		AllowCredentials: true,                    // 是否允许带上用户认证信息（如：cookie），对应响应头：Access-Control-Allow-Credentials
		AllowOriginFunc: func(origin string) bool {
			if strings.HasPrefix(origin, "http://localhost") { // 开发环境用
				return true
			}
			return strings.Contains(origin, "company.domain.com") // 线上环境用
		},
		MaxAge: 12 * time.Hour, // preflity请求有效期，可以调小一点对应响应头：Access-Control-Max-Age
	}))

	redisClient := redis.NewClient(&redis.Options{
		Addr: config.Config.Redis.Addr,
	})

	// 基于redis实现滑动窗口限流。1分钟100次针对IP限流
	server.Use(ratelimit.NewBuilder(redisClient, time.Minute, 100).Build())

	// session存储使用的存储引擎: 内存存储
	//store := cookie.NewStore([]byte("secret"))

	// session存储使用的存储引擎: redis存储
	//store, err := redis.NewStore(16, "tcp", "localhost:6379", "", []byte("dqC2oDoZ2noDoZ2n"), []byte("jW5FOm21NZoDoZ2n"))
	//if err != nil {
	//	panic(err)
	//}

	store := memstore.NewStore([]byte("dqC2oDoZ2noDoZ2n"), []byte("jW5FOm21NZoDoZ2n"))

	server.Use(sessions.Sessions("mysession", store)) // 设置到cookie的name和value

	server.Use(middleware.NewLoginJWTMiddlewareBuilder().
		IgnorePaths("/users/signup").
		IgnorePaths("/users/login").
		Build())
	return server
}
