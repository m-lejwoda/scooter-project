package user

import (
	"context"
	"fmt"
	"time"

	"github.com/redis/go-redis/v9"
)

type TokenStorage struct {
	client *redis.Client
}

func NewUserTokenStorage(client *redis.Client) *TokenStorage {
	return &TokenStorage{client: client}
}

func (t *TokenStorage) GetAccessToken(ctx context.Context, name string) string {
	accessStr := "access_token_"
	key := accessStr + name
	val, err := t.client.Get(ctx, key).Result()
	if err != nil {
		fmt.Printf("Can't get Access Token")
	}
	return val
}

func (t *TokenStorage) GetRefreshToken(ctx context.Context, name string) string {
	refreshStr := "refresh_token_"
	key := refreshStr + name
	val, err := t.client.Get(ctx, key).Result()
	if err != nil {
		fmt.Printf("Can't get refreshToken")
	}
	return val
}

func (t *TokenStorage) DeleteAccessToken(ctx context.Context, name string) {
	accessStr := "access_token_"
	key := accessStr + name
	t.client.Del(ctx, key)
}

func (t *TokenStorage) DeleteRefreshToken(ctx context.Context, name string) {
	refreshStr := "refresh_token_"
	key := refreshStr + name
	t.client.Del(ctx, key)
}

func (t *TokenStorage) CreateAccessTokenBasedOnRefresh(ctx context.Context, name string, accessToken string, refreshToken string) {
	accessStr := "access_token_"
	refreshStr := "refresh_token_"
	accessKey := accessStr + name
	refreshKey := refreshStr + name
	t.client.Set(ctx, accessKey, accessToken, time.Hour*24)
	t.client.Set(ctx, refreshKey, refreshToken, time.Hour*24*30)
}
