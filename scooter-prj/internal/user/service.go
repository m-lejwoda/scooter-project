package user

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type UserService struct {
	repo UserRepository
}

func NewUserService(repo UserRepository) *UserService {
	return &UserService{repo: repo}
}

func (u *UserService) Login(ctx context.Context, user UserLogin) (*UserResponse, error) {
	dbUser, err := u.repo.GetByUsername(ctx, user.Username)
	if err != nil {
		fmt.Println(err)
	}
	var (
		key                []byte
		refreshToken       *jwt.Token
		accessToken        *jwt.Token
		signedRefreshToken string
		signedAccessToken  string
	)
	key = []byte(os.Getenv("JWT_SECRET_KEY"))
	accessToken = jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id":  dbUser.ID,
		"username": dbUser.Username,
		"exp":      time.Now().Add(time.Hour * 24).Unix(),
	})
	signedAccessToken, _ = accessToken.SignedString(key)
	refreshToken = jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id":  dbUser.ID,
		"username": dbUser.Username,
		"exp":      time.Now().Add(time.Hour * 24 * 30).Unix(),
	})
	signedRefreshToken, _ = refreshToken.SignedString(key)
	fmt.Println(signedRefreshToken)
	fmt.Println(signedAccessToken)

	respUser := &UserResponse{
		ID:       dbUser.ID,
		Username: dbUser.Username,
	}
	return respUser, nil
}

func (u *UserService) Register(ctx context.Context, user UserRegister) (*UserResponse, error) {
	createdUser, err := u.repo.Save(ctx, &user)
	if err != nil {
		fmt.Println(err)
	}
	respUser := &UserResponse{
		ID:       createdUser.ID,
		Username: createdUser.Username,
	}
	return respUser, err
}
