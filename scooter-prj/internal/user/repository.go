// Database interface mostly interfaces
package user

import (
	"context"
)

type UserRepository interface {
	getByID(ctx context.Context, id string) (*User, error)
	Save(ctx context.Context, user *User) error
}
