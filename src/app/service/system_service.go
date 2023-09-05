package service

import (
	"net/http"

	"github.com/kallepan/go-backend/app/constant"
	"github.com/kallepan/go-backend/app/pkg"
	"github.com/kallepan/go-backend/app/repository"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

type SystemService interface {
	GetPing(ctx *gin.Context)
}

type SystemServiceImpl struct {
	systemRepository repository.SystemRepository
}

func (s SystemServiceImpl) GetPing(ctx *gin.Context) {
	defer pkg.PanicHandler(ctx)

	data, err := s.systemRepository.GetVersion()
	if err != nil {
		log.Error("Got error when get version: ", err)
		pkg.PanicException(constant.DataNotFound)
	}

	ctx.JSON(http.StatusOK, pkg.BuildResponse(constant.Success, data))
}

func SystemServiceInit(systemRepository repository.SystemRepository) *SystemServiceImpl {
	return &SystemServiceImpl{
		systemRepository: systemRepository,
	}
}
