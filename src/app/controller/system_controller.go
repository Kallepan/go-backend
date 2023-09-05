package controller

import (
	"github.com/kallepan/go-backend/app/service"

	"github.com/gin-gonic/gin"
)

type SystemController interface {
	GetPing(c *gin.Context)
}

type SystemControllerImpl struct {
	svc service.SystemService
}

func (ctrl SystemControllerImpl) GetPing(c *gin.Context) {
	ctrl.svc.GetPing(c)
}

func SystemControllerInit(systemService service.SystemService) *SystemControllerImpl {
	return &SystemControllerImpl{
		svc: systemService,
	}
}
