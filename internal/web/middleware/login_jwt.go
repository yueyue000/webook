package middleware

import (
	"encoding/gob"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/yueyue000/webook/internal/web"
	"log"
	"net/http"
	"time"
)

// LoginJWTMiddleware JWT登录校验
type LoginJWTMiddleware struct {
	paths []string
}

func NewLoginJWTMiddlewareBuilder() *LoginJWTMiddleware {
	return &LoginJWTMiddleware{}
}

func (l *LoginJWTMiddleware) IgnorePaths(path string) *LoginJWTMiddleware {
	l.paths = append(l.paths, path)
	return l
}

// Build builder设计模式, 方便后续扩展
func (l *LoginJWTMiddleware) Build() gin.HandlerFunc {
	gob.Register(time.Now()) // 当编解码中有一个字段是interface{}的时候，需要对interface{}可能产生的类型进行注册
	return func(ctx *gin.Context) {
		// 不需要session校验
		for _, path := range l.paths {
			if ctx.Request.URL.Path == path {
				return
			}
		}
		tokenStr := ctx.GetHeader("X-Jwt-Token")
		if tokenStr == "" {
			ctx.AbortWithStatus(http.StatusUnauthorized)
			return
		}
		claims := &web.UserClaims{}
		token, err := jwt.ParseWithClaims(tokenStr, claims, func(token *jwt.Token) (interface{}, error) {
			return []byte("moyn8y9abnd7q4zkq2m73yw8tu9j5ixm"), nil
		})
		if err != nil {
			ctx.AbortWithStatus(http.StatusUnauthorized)
			return
		}
		// token校验未通过
		if token == nil || !token.Valid || claims.Uid == 0 {
			ctx.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		// 固定间隔时间刷新token，旧的未过期的token也是可以用的
		now := time.Now()
		timeSub := claims.ExpiresAt.Sub(now) < time.Second*20 // 每20s刷新一次token
		if timeSub {
			claims.ExpiresAt = jwt.NewNumericDate(time.Now().Add(time.Minute))
			tokenStr, err := token.SignedString([]byte("moyn8y9abnd7q4zkq2m73yw8tu9j5ixm")) // TODO 如果用户一直用同一个token的话，在token过期之前每次请求都会生成新token，是否会有性能问题，是否需要加风控逻辑。
			if err != nil {
				log.Println("jwt 续约失败", err) // 续约失败不影响后续逻辑，不return
			}
			ctx.Header("X-Jwt-Token", tokenStr)
		}

		ctx.Set("claims", claims)
	}
}
