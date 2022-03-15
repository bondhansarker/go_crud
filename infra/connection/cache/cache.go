package cache

import (
	"context"
	"encoding/json"
	"strconv"
	"time"

	"demo/app/utils/methods"
	"demo/infra/errors"
)

const KeyPrefix = "demo:"

func (cc RedisCacheClient) Set(ctx context.Context, key string, value interface{}, ttl int) error {

	serializedValue, err := json.Marshal(value)
	if err != nil {
		return err
	}

	return cc.Redis.Set(ctx, KeyPrefix+key, string(serializedValue), time.Duration(ttl)*time.Second).Err()
}

func (cc RedisCacheClient) Get(ctx context.Context, key string) (string, error) {
	if methods.IsEmpty(key) {
		return "", errors.ErrEmptyRedisKeyValue
	}

	return cc.Redis.Get(ctx, key).Result()
}

func (cc RedisCacheClient) GetInt(ctx context.Context, key string) (int, error) {
	if methods.IsEmpty(key) {
		return 0, errors.ErrEmptyRedisKeyValue
	}

	str, err := cc.Redis.Get(ctx, key).Result()
	if err != nil {
		return 0, err
	}

	return strconv.Atoi(str)
}

func (cc RedisCacheClient) GetStruct(ctx context.Context, key string, outputStruct interface{}) error {
	if methods.IsEmpty(key) {
		return errors.ErrEmptyRedisKeyValue
	}

	serializedValue, err := cc.Redis.Get(ctx, key).Result()
	if err != nil {
		return err
	}

	if err := json.Unmarshal([]byte(serializedValue), &outputStruct); err != nil {
		return err
	}

	return nil
}

func (cc RedisCacheClient) Del(ctx context.Context, keys ...string) error {
	return cc.Redis.Del(ctx, keys...).Err()
}

func (cc RedisCacheClient) DelPattern(ctx context.Context, pattern string) error {
	iter := cc.Redis.Scan(ctx, 0, pattern, 0).Iterator()

	for iter.Next(ctx) {
		err := cc.Redis.Del(ctx, iter.Val()).Err()
		if err != nil {
			return err
		}
	}

	if err := iter.Err(); err != nil {
		return err
	}

	return nil
}
