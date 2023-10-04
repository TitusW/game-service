package token

import "github.com/redis/go-redis/v9"

type TokenModule struct {
	redis *redis.Client
}

func New(redis *redis.Client) TokenModule {
	return TokenModule{
		redis: redis,
	}
}
