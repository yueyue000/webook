package middleware

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"net/http"
)

type LoginMiddleware struct {
	paths []string
}

func NewLoginMiddlewareBuilder() *LoginMiddleware {
	return &LoginMiddleware{}
}

func (l *LoginMiddleware) IgnorePaths(path string) *LoginMiddleware {
	l.paths = append(l.paths, path)
	return l
}

// Build builder设计模式, 方便后续扩展
func (l *LoginMiddleware) Build() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// 不需要session校验
		for _, path := range l.paths {
			if ctx.Request.URL.Path == path {
				return
			}
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
		uid, ok := id.(int64)
		if !ok {
			ctx.AbortWithStatus(http.StatusUnauthorized) // 没权限
			return
		}
		ctx.Set("uid", uid)
	}
}
