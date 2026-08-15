package user

import (
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type JWTManager struct {
	tokensTTL TokensTTL
}

func NewJWTManager(tokensTTL TokensTTL) *JWTManager {
	return &JWTManager{
		tokensTTL,
	}
}

func (j JWTManager) GenerateTokens(userID string, username string) TokenResponse {
	var (
		key                []byte
		refreshToken       *jwt.Token
		accessToken        *jwt.Token
		signedRefreshToken string
		signedAccessToken  string
	)

	key = []byte(os.Getenv("JWT_SECRET_KEY"))
	accessTokenExpiration := time.Now().Add(j.tokensTTL.AccessTTL).Unix()
	refreshTokenExpiration := time.Now().Add(j.tokensTTL.RefreshTTL).Unix()
	accessToken = jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id":  userID,
		"username": username,
		"exp":      accessTokenExpiration,
	})
	signedAccessToken, _ = accessToken.SignedString(key)
	refreshToken = jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id":  userID,
		"username": username,
		"exp":      refreshTokenExpiration,
	})
	signedRefreshToken, _ = refreshToken.SignedString(key)

	return TokenResponse{
		AccessToken:            signedAccessToken,
		RefreshToken:           signedRefreshToken,
		AccessTokenExpiration:  accessTokenExpiration,
		RefreshTokenExpiration: refreshTokenExpiration,
	}
}
