package repository

import (
	"github.com/gin-gonic/gin"
	"github.com/kallepan/go-backend/app/domain/dao"

	log "github.com/sirupsen/logrus"
)

type UserRepository interface {
	FindAllUsers(ctx *gin.Context) ([]dao.User, error)
	FindUserByID(ctx *gin.Context, userID int) (dao.User, error)
}

type UserRepositoryImpl struct {
}

func (u UserRepositoryImpl) FindUserByID(ctx *gin.Context, userID int) (dao.User, error) {
	var user dao.User

	log.Info("FindUserByID: ", userID)

	return user, nil
}

func (u UserRepositoryImpl) FindAllUsers(ctx *gin.Context) ([]dao.User, error) {
	var users []dao.User

	log.Info("FindAllUsers: ")
	log.Info("It is not implemented yet")

	return users, nil
}

// add ldap or db connection here
func UserRepositoryInit() *UserRepositoryImpl {
	return &UserRepositoryImpl{}
}
