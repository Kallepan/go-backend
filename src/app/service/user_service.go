package service

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/kallepan/go-backend/app/constant"
	"github.com/kallepan/go-backend/app/pkg"
	"github.com/kallepan/go-backend/app/repository"
	log "github.com/sirupsen/logrus"
)

type UserService interface {
	GetAllUser(ctx *gin.Context)
	GetUserById(ctx *gin.Context)
}

type UserServiceImpl struct {
	userRepository repository.UserRepository
}

func (u UserServiceImpl) GetUserById(ctx *gin.Context) {
	defer pkg.PanicHandler(ctx)

	log.Info("Start get user by id")
	userID, _ := strconv.Atoi(ctx.Param("userID"))

	data, err := u.userRepository.FindUserByID(userID)
	if err != nil {
		log.Error(err)
		pkg.PanicException(constant.UnknownError)
	}

	ctx.JSON(200, pkg.BuildResponse(constant.Success, data))
}

func (u UserServiceImpl) GetAllUser(ctx *gin.Context) {
	defer pkg.PanicHandler(ctx)

	log.Info("Start get all user")
	data, err := u.userRepository.FindAllUsers()
	if err != nil {
		log.Error(err)
		pkg.PanicException(constant.UnknownError)
	}
	ctx.JSON(200, pkg.BuildResponse(constant.Success, data))
}

func UserServiceInit(userRepository repository.UserRepository) *UserServiceImpl {
	return &UserServiceImpl{
		userRepository: userRepository,
	}
}
