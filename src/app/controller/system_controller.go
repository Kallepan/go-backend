package controller

import (
	"github.com/kallepan/go-backend/app/service"

	"github.com/gin-gonic/gin"
)

type SystemController interface {
	GetPing(ctx *gin.Context)
}

type SystemControllerImpl struct {
	svc service.SystemService
}

func (ctrl SystemControllerImpl) GetPing(ctx *gin.Context) {
	ctrl.svc.GetPing(ctx)
}

func SystemControllerInit(systemService service.SystemService) *SystemControllerImpl {
	return &SystemControllerImpl{
		svc: systemService,
	}
}
