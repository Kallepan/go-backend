// go:build wireinject
//go:build wireinject
// +build wireinject

package config

import (
	"github.com/kallepan/go-backend/app/controller"
	"github.com/kallepan/go-backend/app/repository"
	"github.com/kallepan/go-backend/app/service"

	"github.com/google/wire"
)

var systemSvcSet = wire.NewSet(service.SystemServiceInit,
	wire.Bind(new(service.SystemService), new(*service.SystemServiceImpl)),
)

var systemCtrlrSet = wire.NewSet(controller.SystemControllerInit,
	wire.Bind(new(controller.SystemController), new(*controller.SystemControllerImpl)),
)

var systemRepoSet = wire.NewSet(repository.SystemRepositoryInit,
	wire.Bind(new(repository.SystemRepository), new(*repository.SystemRepositoryImpl)),
)

func Init() *Initialization {
	wire.Build(
		NewInitialization,
		systemRepoSet,
		systemSvcSet,
		systemCtrlrSet,
	)
	return nil
}
