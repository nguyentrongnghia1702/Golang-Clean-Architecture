package config

import (
	"log"

	"github.com/spf13/viper"
)

var config *viper.Viper

var Appconfig *viper.Viper

func Init() {
	var err error
	config = viper.New()
	config.SetConfigName("config")
	config.AddConfigPath("./mcs-common/config")
	err = config.ReadInConfig()
	if err != nil {
		log.Fatal("Fatal error config file: %s \n", err)
	}

}

func GetConfig() *viper.Viper {
	return config
}
