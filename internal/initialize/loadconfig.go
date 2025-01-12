package initialize

import (
	"fmt"

	"github.com/spf13/viper"
	"github.com/thanhminhdev2000/ecommerce-backend-go/global"
)

func LoadConfig() {
	viper := viper.New()
	viper.AddConfigPath("config")
	viper.SetConfigName("local")
	viper.SetConfigType("yaml")

	// read configuration
	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("fatal error config file: %w", err))
	}

	if err := viper.Unmarshal(&global.Config); err != nil {
		panic(fmt.Errorf("fatal error config file: %w", err))
	}
}