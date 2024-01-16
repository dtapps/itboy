package itboy

import (
	"go.dtapp.net/golog"
)

type Client struct {
	gormLog struct {
		status bool           // 状态
		client *golog.ApiGorm // 日志服务
	}
	mongoLog struct {
		status bool            // 状态
		client *golog.ApiMongo // 日志服务
	}
}

func NewClient() (*Client, error) {
	return &Client{}, nil
}
