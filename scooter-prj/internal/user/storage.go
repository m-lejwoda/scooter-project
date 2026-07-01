// Repository implementation
package user

import (
	"context"
	"fmt"

	"scooter-prj/internal/database"
)

type UserStorage struct {
	db *database.DB
}

func NewUserStorage(db *database.DB) *UserStorage {
	return &UserStorage{db: db}
}

func (u *UserStorage) GetByID(ctx context.Context, id string) (*User, error) {
	query := `SELECT id, username, lastname FROM users WHERE id = $1:`
	var user User
	err := u.db.Pool.QueryRow(ctx, query, id).Scan(&user.ID, &user.Username, &user.Lastname)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch user: %w", err)
	}
	return &user, nil
}
