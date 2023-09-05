package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/kallepan/go-backend/app/service"
)

type UserController interface {
	GetAllUsers(ctx *gin.Context)
	GetUserById(ctx *gin.Context)
}

type UserControllerImpl struct {
	svc service.UserService
}

func (u UserControllerImpl) GetAllUsers(ctx *gin.Context) {
	u.svc.GetAllUser(ctx)
}

func (u UserControllerImpl) GetUserById(ctx *gin.Context) {
	u.svc.GetUserById(ctx)
}

func UserControllerInit(userService service.UserService) *UserControllerImpl {
	return &UserControllerImpl{
		svc: userService,
	}
}
