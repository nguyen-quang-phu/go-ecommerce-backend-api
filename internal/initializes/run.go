package initializes

import (
	"fmt"

	"github.com/nguyen-quang-phu/go-ecommerce-backend-api/global"
)

func Run() {
	loadConfig()
	InitLogger()
	InitDatabase()
	InitServiceInterface()
	InitRedis()
	r := InitRouter()

	port := fmt.Sprintf(":%d", global.Config.Server.Port)
	err := r.Run(port)
	if err != nil {
		return
	}
}
