package middleware

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"net/http"
)

type LoginMiddleware struct {
}

func NewLoginMiddlewareBuilder() *LoginMiddleware {
	return &LoginMiddleware{}
}

// Build builder设计模式, 方便后续扩展
func (l *LoginMiddleware) Build() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// 不需要登录校验的接口
		if ctx.Request.URL.Path == "/users/login" ||
			ctx.Request.URL.Path == "/users/signup" {
			return
		}

		session := sessions.Default(ctx)
		if session == nil {
			ctx.AbortWithStatus(http.StatusUnauthorized) // 没权限
			return
		}
		id := session.Get("userID")
		if id == nil {
			ctx.AbortWithStatus(http.StatusUnauthorized) // 没权限
			return
		}
	}
}
