package database

import (
	"fmt"
	"mcs-nghiadeptrai/mcs-common/config"
	"mcs-nghiadeptrai/mcs-common/logger"
	"time"

	"github.com/allegro/bigcache"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	DB          *gorm.DB
	GlobalCache *bigcache.BigCache
)

func Init() {
	logger.LogInfoNoContext("TESSTT!!")

	var err error
	dsn := config.Appconfig.GetString("database.username") + ":" + config.Appconfig.GetString("database.password") + "@tcp(" + config.Appconfig.GetString("database.host") + ":" + config.Appconfig.GetString("database.port") + ")/" + config.Appconfig.GetString("database.name") + "?parseTime=true"
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		logger.LogErrorNoContext(err.Error())
		panic(fmt.Errorf("failed to connect to DB: %w", err))
	}
	GlobalCache, err = bigcache.NewBigCache(bigcache.DefaultConfig(30 * time.Minute))
	if err != nil {
		panic(fmt.Errorf("failed to initialize cahce: %w", err))
	}
}
