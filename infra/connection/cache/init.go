package cache

import (
	"demo/app/domain"
)

var client RedisCacheClient

func NewCacheClient() domain.ICache {
	ConnectRedis()
	return &RedisCacheClient{}
}

func Client() RedisCacheClient {
	return client
}
