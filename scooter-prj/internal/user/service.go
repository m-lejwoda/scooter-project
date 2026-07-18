package user

import (
	"context"
	"fmt"
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
