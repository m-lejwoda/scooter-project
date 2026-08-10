package user

import (
	"context"
	"fmt"

	"github.com/redis/go-redis/v9"
)

type TokenStorage struct {
	client *redis.Client
}

func NewUserTokenStorage(client *redis.Client) *TokenStorage {
	return &TokenStorage{client: client}
}

func (t *TokenStorage) GetAccessToken(ctx context.Context, id string) string {
	accessStr := "access_token_"
	key := accessStr + id
	val, err := t.client.Get(ctx, key).Result()
	if err != nil {
		fmt.Printf("Can't get Access Token")
	}
	return val
}

func (t *TokenStorage) GetRefreshToken(ctx context.Context, id string) string {
	refreshStr := "refresh_token_"
	key := refreshStr + id
	val, err := t.client.Get(ctx, key).Result()
	if err != nil {
		fmt.Printf("Can't get refreshToken")
	}
	return val
}

func (t *TokenStorage) DeleteAccessToken(ctx context.Context, id string) {
}

func (t *TokenStorage) DeleteRefreshToken(ctx context.Context, id string) {
}

func (t *TokenStorage) CreateAccessTokenBasedOnRefresh(ctx context.Context, id string) {
}

func (t *TokenStorage) CreateTokens(ctx context.Context, id string) {
}
