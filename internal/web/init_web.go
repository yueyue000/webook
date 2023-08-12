package web

import "github.com/gin-gonic/gin"

func RegisterRoutes(s *gin.Engine) {
	u := &UserHandler{}
	ug := s.Group("/users")
	ug.POST("/signup", u.SignUp)
	ug.POST("/login", u.Login)
	ug.POST("/edit", u.Edit)
	ug.POST("/profile", u.Profile)
}
