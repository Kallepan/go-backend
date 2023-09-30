// go:build wireinject
//go:build wireinject
// +build wireinject

package config

import (
	"github.com/kallepan/go-backend/app/controller"
	"github.com/kallepan/go-backend/app/repository"
	"github.com/kallepan/go-backend/app/service"
	"github.com/kallepan/go-backend/drivers"

	"github.com/google/wire"
)

// Set of providers for initialization

/* Database */
var db = wire.NewSet(drivers.ConnectToDB)

/* system */
var (
	systemRepoSet = wire.NewSet(repository.SystemRepositoryInit,
		wire.Bind(new(repository.SystemRepository), new(*repository.SystemRepositoryImpl)),
	)
	systemSvcSet = wire.NewSet(service.SystemServiceInit,
		wire.Bind(new(service.SystemService), new(*service.SystemServiceImpl)),
	)
	systemCtrlrSet = wire.NewSet(controller.SystemControllerInit,
		wire.Bind(new(controller.SystemController), new(*controller.SystemControllerImpl)),
	)
)

/* user */
var (
	userRepoSet = wire.NewSet(repository.UserRepositoryInit,
		wire.Bind(new(repository.UserRepository), new(*repository.UserRepositoryImpl)),
	)
	userSvcSet = wire.NewSet(service.UserServiceInit,
		wire.Bind(new(service.UserService), new(*service.UserServiceImpl)),
	)
	userCtrlrSet = wire.NewSet(controller.UserControllerInit,
		wire.Bind(new(controller.UserController), new(*controller.UserControllerImpl)),
	)
)

func Init() *Initialization {
	wire.Build(
		NewInitialization,
		db,
		systemRepoSet,
		systemSvcSet,
		systemCtrlrSet,
		userRepoSet,
		userSvcSet,
		userCtrlrSet,
	)
	return nil
}
