package main

import (
	_ "github.com/jinzhu/gorm/dialects/postgres"

	_ "oasis-be/database"
	_ "oasis-be/minio"
	_ "oasis-be/redis"
	routers "oasis-be/routers"

	"github.com/astaxie/beego"
	"github.com/gin-gonic/gin"
)

func main() {
	if beego.BConfig.RunMode == "dev" {
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	}
	// beego.Run()

	runMode := beego.AppConfig.DefaultString("gin.mode", "debug")
	serverPort := beego.AppConfig.DefaultString("httpport", "8080")

	gin.SetMode(runMode)
	routersInit := routers.InitRouter()
	routersInit.Run(":" + serverPort)

}
