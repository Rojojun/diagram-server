package database

import (
	"context"
	"fmt"
)

type Connector interface {
	Connect(ctx context.Context) error
	Disconnect(ctx context.Context) error
	Ping(ctx context.Context) error
	Client() any
}
type ConnectType string

const (
	MongoDB ConnectType = "mongodb"
)

func NewConnector(cfg Config) (Connector, error) {
	switch cfg.Type {
	case MongoDB:
		return NewMongoConnector(cfg), nil
	default:
		return nil, fmt.Errorf("unsupported database type: %s", cfg.Type)
	}
}
