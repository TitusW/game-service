package token

import (
	"context"
	"fmt"
	"time"
)

func (m TokenModule) SetUserToken(ctx context.Context, userKsuid string, token string) error {
	key := userKsuid + ":" + token
	if err := m.redis.Set(ctx, key, 0, 15*time.Minute).Err(); err != nil {
		return err
	}

	return nil
}

func (m TokenModule) DeleteUserToken(ctx context.Context, userKsuid string, token string) error {
	key := fmt.Sprintf("%s:%s", userKsuid, token)

	result := m.redis.Del(ctx, key)

	if err := result.Err(); err != nil {
		return err
	}

	if result.Val() < 1 {
		return fmt.Errorf("User token does not exist")
	}

	return nil
}

func (m TokenModule) ScanUserTokens(ctx context.Context, userKsuid string, token string) ([]string, error) {
	var keys []string
	var match string

	switch {
	case token != "":
		match = fmt.Sprintf("%s:%s", userKsuid, token)
	default:
		match = fmt.Sprintf("%s:*", userKsuid)
	}
	iter := m.redis.Scan(ctx, 0, match, 0).Iterator()

	if err := iter.Err(); err != nil {
		return []string{}, err
	}

	for iter.Next(ctx) {
		keys = append(keys, iter.Val())
	}

	return keys, nil
}
