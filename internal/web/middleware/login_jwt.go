package middleware

import (
	"encoding/gob"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
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
		session := sessions.Default(ctx)
		if session == nil {
			ctx.AbortWithStatus(http.StatusUnauthorized) // 没权限
			return
		}
		id := session.Get("uid")
		if id == nil {
			ctx.AbortWithStatus(http.StatusUnauthorized) // 没权限
			return
		}
		uid, ok := id.(int64)
		if !ok {
			ctx.AbortWithStatus(http.StatusUnauthorized) // 没权限
			return
		}

		// 将uid写入param方便后面使用
		ctx.AddParam("uid", strconv.FormatInt(uid, 10))

		// 刷新session
		updateTime := session.Get("update_time")
		session.Set("uid", id)
		session.Options(sessions.Options{
			MaxAge: 30,
		})
		now := time.Now()
		if updateTime == nil { // 刚登陆，还未刷新过
			session.Set("update_time", now)
			err := session.Save()
			if err != nil {
				panic(err)
			}
			return
		}

		updateTimeVal, ok := updateTime.(time.Time)
		if !ok {
			ctx.AbortWithStatus(http.StatusInternalServerError)
			return
		}
		if now.Sub(updateTimeVal) > time.Second*10 {
			session.Set("update_time", now)
			err := session.Save()
			if err != nil {
				panic(err)
			}
			return
		}
	}
}
