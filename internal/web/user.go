package web

import "github.com/gin-gonic/gin"

// UserHandler 定义所有跟user有关的路由
type UserHandler struct {
}

func (u *UserHandler) SignUp(ctx *gin.Context) {}

func (u *UserHandler) Login(ctx *gin.Context) {}

func (u *UserHandler) Edit(ctx *gin.Context) {}

func (u *UserHandler) Profile(ctx *gin.Context) {}
