package web

import (
	"errors"
	"fmt"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator"
	"github.com/yueyue000/webook/internal/domain"
	"github.com/yueyue000/webook/internal/service"
	"net/http"
)

// UserHandler 定义所有跟user有关的路由
type UserHandler struct {
	svc *service.UserService
}

func (u *UserHandler) RegisterRoutes(s *gin.Engine) {
	ug := s.Group("/users")
	ug.POST("/signup", u.SignUp)
	ug.POST("/login", u.Login)
	ug.POST("/edit", u.Edit)
	ug.GET("/profile", u.Profile)
}

func NewUserHandler(svc *service.UserService) *UserHandler {
	return &UserHandler{
		svc: svc,
	}
}

func (u *UserHandler) SignUp(ctx *gin.Context) {
	type signUpReq struct {
		Email    string `json:"email" validate:"email"`
		Password string `json:"password" validate:"min=6,max=16"`
	}

	var req signUpReq
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

	err = u.svc.SignUp(ctx, domain.User{
		Email:    req.Email,
		Password: req.Password,
	})
	if errors.Is(err, service.ErrUserDuplicateEmail) {
		ctx.String(http.StatusOK, "邮箱冲突")
		return
	}

	if err != nil {
		ctx.String(http.StatusOK, "系统异常")
		return
	}

	ctx.String(http.StatusOK, "注册成功")
	fmt.Printf("req:%+v", req)
}

func (u *UserHandler) Login(ctx *gin.Context) {
	type loginReq struct {
		Email    string `json:"email" validate:"email"`
		Password string `json:"password" validate:"required"`
	}

	var req loginReq
	if err := ctx.Bind(&req); err != nil {
		ctx.JSON(499, gin.H{"msg": err.Error()})
		return
	}
	user, err := u.svc.Login(ctx, domain.User{
		Email:    req.Email,
		Password: req.Password,
	})
	if err == service.ErrInvalidUserOrPassword {
		ctx.String(http.StatusOK, "用户名或密码错误")
		return
	}
	if err != nil {
		ctx.String(http.StatusOK, "系统错误")
		return
	}

	// 登录成功，获取sid, 设置sid
	session := sessions.Default(ctx)
	session.Set("userID", user.ID) // 只能设置一次，不能设置多个
	session.Save()

	ctx.JSON(http.StatusOK, gin.H{
		"msg": "登录成功",
	})
	return
}

func (u *UserHandler) Edit(ctx *gin.Context) {
	type editReq struct {
		Nick     string `json:"nick" validate:"min=0,max=16"`
		Birthday string `json:"birthday" validate:"datetime=2006-01-02"`
	}

	var req editReq
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
	ctx.String(http.StatusOK, "修改成功")
}

func (u *UserHandler) Profile(ctx *gin.Context) {
	ctx.String(http.StatusOK, "这里是资料")
}
