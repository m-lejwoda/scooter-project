// Database interface mostly interfaces
package user

import (
	"context"
)

type UserRepository interface {
	GetByID(ctx context.Context, id string) (*User, error)
	GetByUsername(ctx context.Context, user *UserLogin) (*UserResponse, error)
	Save(ctx context.Context, user *UserRegister) error
}
