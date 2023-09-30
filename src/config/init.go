package config

import (
	"database/sql"

	"github.com/kallepan/go-backend/app/controller"
	"github.com/kallepan/go-backend/app/repository"
	"github.com/kallepan/go-backend/app/service"
)

type Initialization struct {
	db       *sql.DB
	sysRepo  repository.SystemRepository
	sysSvc   service.SystemService
	SysCtrl  controller.SystemController
	userRepo repository.UserRepository
	userSvc  service.UserService
	UserCtrl controller.UserController
}

func NewInitialization(
	db *sql.DB,
	sysRepo repository.SystemRepository,
	sysSvc service.SystemService,
	sysCtrl controller.SystemController,
	userRepo repository.UserRepository,
	userSvc service.UserService,
	userCtrl controller.UserController,
) *Initialization {
	return &Initialization{
		db:       db,
		sysRepo:  sysRepo,
		sysSvc:   sysSvc,
		SysCtrl:  sysCtrl,
		userRepo: userRepo,
		userSvc:  userSvc,
		UserCtrl: userCtrl,
	}
}
