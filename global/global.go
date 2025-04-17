package global

import (
	"github.com/nguyen-quang-phu/go-ecommerce-backend-api/pkg/logger"
	"github.com/nguyen-quang-phu/go-ecommerce-backend-api/pkg/settings"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

var (
	Config settings.Config
	Logger *logger.LoggerZap
	DB     *gorm.DB
	Rdb    *redis.Client
)
