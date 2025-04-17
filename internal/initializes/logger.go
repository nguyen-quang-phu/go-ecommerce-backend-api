package initializes

import (
	"github.com/nguyen-quang-phu/go-ecommerce-backend-api/global"
	"github.com/nguyen-quang-phu/go-ecommerce-backend-api/pkg/logger"
)

func InitLogger() {
	global.Logger = logger.NewLogger(global.Config.Logger)
}
