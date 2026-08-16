package user

import (
	"context"
	"fmt"
	"os"
	"strconv"

	"github.com/golang-jwt/jwt/v5"
)

type UserService struct {
	userRepo   UserRepository
	tokenRepo  TokenRepository
	jwtManager *JWTManager
}

func NewUserService(userRepo UserRepository, tokenRepo TokenRepository, jwtManager *JWTManager) *UserService {
	return &UserService{userRepo: userRepo, tokenRepo: tokenRepo, jwtManager: jwtManager}
}

func (u *UserService) Login(ctx context.Context, user UserLogin) (*UserResponse, error) {
	dbUser, err := u.userRepo.GetByUsername(ctx, user.Username)
	if err != nil {
		fmt.Println(err)
	}
	respUser, err := u.generateUserResponse(ctx, dbUser.ID, dbUser.Username)
	return respUser, err
}

func (u *UserService) Register(ctx context.Context, user UserRegister) (*UserResponse, error) {
	createdUser, err := u.userRepo.Save(ctx, &user)
	if err != nil {
		fmt.Println(err)
	}
	respUser, err := u.generateUserResponse(ctx, createdUser.ID, createdUser.Username)
	return respUser, err
}

func (u *UserService) RefreshToken(ctx context.Context, token string, userID int, username string) (*UserResponse, error) {
	respUser, err := u.generateUserResponse(ctx, userID, username)
	return respUser, err
}

func (u *UserService) generateUserResponse(ctx context.Context, userID int, username string) (*UserResponse, error) {
	userIDStr := strconv.Itoa(userID)

	generatedTokens := u.jwtManager.GenerateTokens(userIDStr, username)

	name := username + userIDStr
	u.tokenRepo.DeleteAccessToken(ctx, name)
	u.tokenRepo.DeleteRefreshToken(ctx, name)
	token, err := u.tokenRepo.CreateTokens(ctx, name, generatedTokens, u.jwtManager.tokensTTL)
	if err != nil {
		return nil, fmt.Errorf("failed to store tokens: %w", err)
	}

	return &UserResponse{
		ID:                     userID,
		Username:               username,
		AccessToken:            token.AccessToken,
		RefreshToken:           token.RefreshToken,
		AccessTokenExpiration:  token.AccessTokenExpiration,
		RefreshTokenExpiration: token.RefreshTokenExpiration,
	}, nil
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
