package main

import (
	"os"

	"github.com/kallepan/go-backend/app/router"
	"github.com/kallepan/go-backend/config"
)

func main() {
	port := os.Getenv("PORT")

	config.InitlLog()

	init := config.Init()
	app := router.Init(init)

	app.Run(":" + port)
}
