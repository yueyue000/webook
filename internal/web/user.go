package web

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator"
	"net/http"
)

// UserHandler 定义所有跟user有关的路由
type UserHandler struct {
}

func (u *UserHandler) SignUp(ctx *gin.Context) {
	type SignUpReq struct {
		Email    string `json:"email" validate:"email"`
		Password string `json:"password" validate:"min=6,max=18"`
	}

	var req SignUpReq
	// Bind方法会根据Content-Type来解析请求参数到req里面，解析出错会直接写回400的错误
	if err := ctx.Bind(&req); err != nil {
		ctx.JSON(499, gin.H{"msg": err.Error()})
		return
	}

	validate := validator.New()
	err := validate.Struct(req)
	if err != nil {
		ctx.JSON(499, gin.H{"msg": err.Error()})
		return
	}

	ctx.String(http.StatusOK, "注册成功")
	fmt.Printf("req:%+v", req)
}

func (u *UserHandler) Login(ctx *gin.Context) {}

func (u *UserHandler) Edit(ctx *gin.Context) {}

func (u *UserHandler) Profile(ctx *gin.Context) {}
