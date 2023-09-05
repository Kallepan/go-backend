package router

import (
	"github.com/kallepan/go-backend/app/middleware"
	"github.com/kallepan/go-backend/config"

	"github.com/gin-gonic/gin"
)

func Init(init *config.Initialization) *gin.Engine {

	router := gin.New()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())
	router.Use(middleware.Cors())

	api := router.Group("/api/v1")
	{
		api.GET("/ping", init.SysCtrl.GetPing)
	}

	return router
}
