package v1

import (
	"v1/internal/service"
	"v1/pkg/app"
	"v1/pkg/err"

	"github.com/gin-gonic/gin"
)

type User struct {
}

func NewUser() User {
	return User{}
}

func (u User) Login(c *gin.Context) {
	param := service.LoginRequest{}
	response := app.NewResponse(c)
	valid, errs := app.BindAndValid(c, &param)
	if !valid {
		response.ErrTo(err.InvalidParams.WithDetails(errs.Errors()...))
		return
	}
	response.To(gin.H{})
}
