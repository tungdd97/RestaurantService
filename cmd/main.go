package main

import (
	"github.com/gin-gonic/gin"
	"restaurant-service/api/route"
	"restaurant-service/bootstrap"
	"time"
)

func main() {

	app := bootstrap.App()

	env := app.Env

	db := app.Mongo.Database(env.DBName)
	defer app.CloseDBConnection()

	timeout := time.Duration(env.ContextTimeout) * time.Second

	appRun := gin.Default()

	route.Setup(env, timeout, db, appRun)

	err := appRun.Run(env.ServerAddress)
	if err != nil {
		return
	}
}
