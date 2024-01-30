package main

import (
	config "mcs-nghiadeptrai/mcs-common/config"
	"mcs-nghiadeptrai/mcs-common/database"
	logger "mcs-nghiadeptrai/mcs-common/logger"
	"mcs-nghiadeptrai/mcs-common/router"
)

func main() {
	config.Init()
	config.Appconfig = config.GetConfig()

	logger.Init()
	logger.LogInfoNoContext("INIT LOGGER SUCCESSFUL")

	database.Init()
	logger.LogInfoNoContext("INIT DATABASE SUCCESSFUL")

	router.Init()
	logger.LogInfoNoContext("INIT ROUTER SUCCESSFUL")

}
