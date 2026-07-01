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

func (u *UserService) Login(ctx context.Context, user User) error {
	err := u.repo.Save(ctx, &user)
	if err != nil {
		fmt.Println(err)
	}
	return nil
}
