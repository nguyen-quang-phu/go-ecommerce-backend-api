package global

import (
	"database/sql"

	"github.com/nguyen-quang-phu/go-ecommerce-backend-api/pkg/logger"
	"github.com/nguyen-quang-phu/go-ecommerce-backend-api/pkg/settings"
	"github.com/redis/go-redis/v9"
)

var (
	Config settings.Config
	Logger *logger.LoggerZap
	DB     *sql.DB
	Rdb    *redis.Client
)
