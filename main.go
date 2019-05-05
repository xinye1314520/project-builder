package main

import (
	"github.com/gin-gonic/gin"
	"github.com/robfig/cron"
	"project-builder/util"
	"project-builder/route"
	"fmt"
	"project-builder/db"
)

func main() {

	gin.SetMode(gin.DebugMode)

	routeEngine := route.SetupRouter()

	//定时检查数据库连接状态状态
	checkDbCron := cron.New()
	checkDbCron.AddFunc("*/10 * * * * *", db.CheckDB)
	//checkDbCron.Start()

	port := util.Config.Section("system").Key("http.port").String()
	routeEngine.Run(":" + port)
	fmt.Println("project run ... ")
}
