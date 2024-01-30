package router

import (
	"log"
	"mcs-nghiadeptrai/mcs-common/config"
	"mcs-nghiadeptrai/mcs-common/database"
	"mcs-nghiadeptrai/mcs-common/middleware"
	"mcs-nghiadeptrai/mcs-todo-item/api/route"
	"time"

	"github.com/fatih/color"
	"github.com/gin-gonic/gin"
)

func Init() {
	router := NewRouter()

	router.Run(config.Appconfig.GetString("Server.port"))
}

func NewRouter() *gin.Engine {
	router := gin.New()
	gin.ForceConsoleColor()
	green := color.New(color.FgGreen).SprintFunc()
	gin.DebugPrintRouteFunc = func(httpMethod, absolutePath, handlerName string, nuHandlers int) {
		log.Printf("%v %v %v handlers %v - Registered routing %v \n", green("INFO"), handlerName, nuHandlers, httpMethod, absolutePath)
	}

	resource := router.Group(config.Appconfig.GetString("Server.context-path"))
	resource.Use(middleware.LogRequestInfo())

	timeout := time.Duration(config.Appconfig.GetInt("Server.context-timeout")) * time.Second

	route.NewTodoItemRouter(timeout, database.DB, resource)
	return router
}
