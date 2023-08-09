package web

import "github.com/gin-gonic/gin"

func RegisterRoutes(s *gin.Engine) {
	u := &UserHandler{}
	s.Group("/users")
	s.POST("/signup", u.SignUp)
	s.POST("/login", u.Login)
	s.POST("/edit", u.Edit)
	s.POST("/profile", u.Profile)
}
