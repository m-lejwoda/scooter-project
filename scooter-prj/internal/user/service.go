package user

import (
	"context"
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type UserService struct {
	userRepo  UserRepository
	tokenRepo TokenRepository
}

func NewUserService(userRepo UserRepository, tokenRepo TokenRepository) *UserService {
	return &UserService{userRepo: userRepo, tokenRepo: tokenRepo}
}

func (u *UserService) Login(ctx context.Context, user UserLogin) (*UserResponse, error) {
	dbUser, err := u.userRepo.GetByUsername(ctx, user.Username)
	if err != nil {
		fmt.Println(err)
	}
	accessTTL := time.Hour * 24
	refreshTTL := time.Hour * 24 * 30
	tokensTTL := TokensTTL{
		AccessTTL:  accessTTL,
		RefreshTTL: refreshTTL,
	}
	j := NewJWTManager(tokensTTL)
	generatedTokens := j.GenerateTokens(strconv.Itoa(dbUser.ID), dbUser.Username)
	name := dbUser.Username + strconv.Itoa(dbUser.ID)
	// Delete old tokens
	u.tokenRepo.DeleteAccessToken(ctx, name)
	u.tokenRepo.DeleteRefreshToken(ctx, name)
	// CreateFreshTokens
	token, err := u.tokenRepo.CreateTokens(ctx, name, generatedTokens, tokensTTL)
	if err != nil {
		fmt.Println(err)
	}
	respUser := &UserResponse{
		ID:                     dbUser.ID,
		Username:               dbUser.Username,
		AccessToken:            token.AccessToken,
		RefreshToken:           token.RefreshToken,
		AccessTokenExpiration:  token.AccessTokenExpiration,
		RefreshTokenExpiration: token.RefreshTokenExpiration,
	}
	return respUser, nil
}

func (u *UserService) Register(ctx context.Context, user UserRegister) (*UserResponse, error) {
	createdUser, err := u.userRepo.Save(ctx, &user)
	if err != nil {
		fmt.Println(err)
	}
	respUser := &UserResponse{
		ID:       createdUser.ID,
		Username: createdUser.Username,
	}
	return respUser, err
}

func (u *UserService) RefreshToken(ctx context.Context, token string) {
	u.tokenRepo.GetRefreshToken(ctx, token)
}

func VerifyToken(tokenString string) (jwt.MapClaims, error) {
	key := []byte(os.Getenv("JWT_SECRET_KEY"))
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpeceted algorithm signing", token.Header["alg"])
		}
		return key, nil
	})
	if err != nil {
		return nil, err
	}
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return claims, nil
	}
	return nil, fmt.Errorf("Wrong token")
}
