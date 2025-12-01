package database

import "time"

type Config struct {
	Type     ConnectType
	Uri      string
	Database string
	Pool     PoolConfig
}

type PoolConfig struct {
	MinSize     uint64
	MaxSize     uint64
	MaxIdleTime time.Duration
	MaxLifetime time.Duration
}

func DefaultPoolConfig() PoolConfig {
	return PoolConfig{
		MinSize:     5,
		MaxSize:     100,
		MaxIdleTime: 30 * time.Second,
		MaxLifetime: 1 * time.Hour,
	}
}
