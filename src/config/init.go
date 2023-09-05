package config

import (
	"github.com/kallepan/go-backend/app/controller"
	"github.com/kallepan/go-backend/app/repository"
	"github.com/kallepan/go-backend/app/service"
)

type Initialization struct {
	sysRepo repository.SystemRepository
	sysSvc  service.SystemService
	SysCtrl controller.SystemController
}

func NewInitialization(
	sysRepo repository.SystemRepository,
	sysSvc service.SystemService,
	sysCtrl controller.SystemController,
) *Initialization {
	return &Initialization{
		sysRepo: sysRepo,
		sysSvc:  sysSvc,
		SysCtrl: sysCtrl,
	}
}
