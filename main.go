package main

import (
	"github.com/martin-flower/api/app"
	"github.com/martin-flower/api/app/config"
)

func main() {
	config := config.GetConfig()

	app := &app.App{}
	app.InitializePostgres(config)
	app.Run(":3000")
}
