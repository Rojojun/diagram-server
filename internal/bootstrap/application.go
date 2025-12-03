package bootstrap

import (
	"context"
	"diagram-server/internal/database"
	"os"
	"strconv"
	"time"
)

func InitDatabase(ctx context.Context) (database.Connector, error) {
	cfg := database.Config{
		Type:     database.MongoDB,
		Uri:      "mongodb://localhost:27017",
		Database: "diagram",
		Pool: database.PoolConfig{
			MinSize:     getEnvUint64("DB_POOL_MIN", 5),
			MaxSize:     getEnvUint64("DB_POOL_MAX", 100),
			MaxIdleTime: getEnvDuration("DB_POOL_IDLE_TIME", 30*time.Second),
		},
	}

	conn, err := database.NewConnector(cfg)
	if err != nil {
		return nil, err
	}

	if err := conn.Connect(ctx); err != nil {
		return nil, err
	}

	if err := conn.Ping(ctx); err != nil {
		return nil, err
	}

	return conn, nil
}

func getEnvUint64(key string, defaultVal uint64) uint64 {
	if v := os.Getenv(key); v != "" {
		if parsed, err := strconv.ParseUint(v, 10, 64); err == nil {
			return parsed
		}
	}
	return defaultVal
}

func getEnvDuration(key string, defaultVal time.Duration) time.Duration {
	if v := os.Getenv(key); v != "" {
		if parsed, err := time.ParseDuration(v); err == nil {
			return parsed
		}
	}
	return defaultVal
}
