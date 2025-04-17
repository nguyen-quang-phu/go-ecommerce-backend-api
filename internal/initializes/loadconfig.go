package initializes

import (
	"fmt"

	"github.com/nguyen-quang-phu/go-ecommerce-backend-api/global"
	"github.com/spf13/viper"
)

func loadConfig() {
	viper := viper.New()
	viper.AddConfigPath("./config/")
	viper.SetConfigName("local")
	viper.SetConfigType("yaml")

	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("failed to read configuration %w \n", err))
	}

	if err := viper.Unmarshal(&global.Config); err != nil {
		fmt.Printf("Unable to decode configuration %v", err)
	}

	fmt.Println("load config",global.Config)
}
