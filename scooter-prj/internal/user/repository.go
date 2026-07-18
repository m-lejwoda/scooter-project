// Database interface mostly interfaces
package user

import (
	"context"
)

type UserRepository interface {
	GetByID(ctx context.Context, id string) (*User, error)
	GetByUsername(ctx context.Context, username string) (*User, error)
	Save(ctx context.Context, user *UserRegister) (*User, error)
}
