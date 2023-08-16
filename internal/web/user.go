package web

import (
	"fmt"
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

func (c *UserHandler) RegisterRoutes(s *gin.Engine) {
	ug := s.Group("/users")
	ug.POST("/signup", c.SignUp)
	ug.POST("/login", c.Login)
	ug.POST("/edit", c.Edit)
	ug.POST("/profile", c.Profile)
}

func NewUserHandler(svc *service.UserService) *UserHandler {
	return &UserHandler{
		svc: svc,
	}
}

func (c *UserHandler) SignUp(ctx *gin.Context) {
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

	err = c.svc.SignUp(ctx, domain.User{
		Email:    req.Email,
		Password: req.Password,
	})
	if err != nil {
		ctx.String(http.StatusOK, "系统异常")
		return
	}

	ctx.String(http.StatusOK, "注册成功")
	fmt.Printf("req:%+v", req)
}

func (c *UserHandler) Login(ctx *gin.Context) {}

func (c *UserHandler) Edit(ctx *gin.Context) {
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

func (c *UserHandler) Profile(ctx *gin.Context) {

}
