package main

import (
	"context"
	"os"

	"github.com/kallepan/go-backend/app/router"
	"github.com/kallepan/go-backend/config"
	"github.com/kallepan/go-backend/drivers"
)

func main() {
	ctx := context.Background()
	port := os.Getenv("PORT")

	config.InitlLog()
	drivers.Init(ctx)

	init := config.Init()
	app := router.Init(init)

	app.Run(":" + port)
}
