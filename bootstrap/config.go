package bootstrap

import (
	"log"

	"github.com/spf13/viper"
)

func InitConfig() *viper.Viper {
	config := viper.New()

	config.SetConfigFile(`config.yaml`)
	err := config.ReadInConfig()
	if err != nil {
		panic("Cant Find File config.yaml")
	}

	if config.GetBool(`debug`) {
		log.Println("Service RUN on DEBUG mode")
	}

	return config
}
