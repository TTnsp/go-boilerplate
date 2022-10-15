package main

import (
	"github.com/ttnsp/go-boilerplate/configuration"
	"github.com/ttnsp/go-boilerplate/repositories"

	"github.com/gin-gonic/gin"
)

func main() {

	configuration.LoadConfig()

	repositories.ConnectDatabase(configuration.App.Database.Host,
		configuration.App.Database.Port,
		configuration.App.Database.Dbname,
		configuration.App.Database.User,
		configuration.App.Database.Password)

	r := gin.Default()
	CreateRoute(r)
	r.Run()
}
