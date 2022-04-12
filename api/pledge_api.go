package main

import (
	"github.com/gin-gonic/gin"
	"pledge-backend/api/middlewares"
	"pledge-backend/api/models/kucoin"
	"pledge-backend/api/models/ws"
	"pledge-backend/api/routes"
	"pledge-backend/api/static"
	"pledge-backend/api/validate"
	"pledge-backend/config"
	"pledge-backend/db"
)

func main() {

	//init mysql
	db.InitMysql()

	//init redis
	db.InitRedis()

	//gin bind go-playground-validator
	validate.BindingValidator()

	// websocket server
	go ws.StartServer()

	// get plgr price from kucoin-exchange
	go kucoin.GetExchangePrice()

	// gin start
	gin.SetMode(gin.ReleaseMode)
	app := gin.Default()
	staticPath := static.GetCurrentAbPathByCaller()
	app.Static("/storage/", staticPath)
	app.Use(middlewares.Cors()) // 「 Cross domain Middleware 」
	routes.InitRoute(app)
	_ = app.Run(":" + config.Config.Env.Port)

}

/*
 If you change the version, you need to modify the following files'
 config/init.go
*/
