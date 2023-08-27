package middleware

import (
	"encoding/gob"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/yueyue000/webook/internal/web"
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
		if token == nil || !token.Valid || claims.Uid == 0 {
			ctx.AbortWithStatus(http.StatusUnauthorized)
			return
		}
		ctx.Set("claims", claims)
	}
}
