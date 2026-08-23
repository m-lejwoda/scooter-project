// Repository implementation
package user

import (
	"context"
	"fmt"

	"scooter-prj/internal/database"

	"golang.org/x/crypto/bcrypt"
)

type UserStorage struct {
	db *database.DB
}

func NewUserStorage(db *database.DB) *UserStorage {
	return &UserStorage{db: db}
}

func (u *UserStorage) GetByID(ctx context.Context, id string) (*User, error) {
	query := `SELECT id, username, lastname FROM users WHERE id = $1`
	var user User
	err := u.db.Pool.QueryRow(ctx, query, id).Scan(&user.ID, &user.Username, &user.Lastname)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch user: %w", err)
	}
	return &user, nil
}

func (u *UserStorage) GetByUsername(ctx context.Context, username string) (*User, error) {
	query := `SELECT id, username, lastname, password, created_at FROM users WHERE username = $1`
	var user User
	err := u.db.Pool.QueryRow(ctx, query, username).Scan(
		&user.ID,
		&user.Username,
		&user.Lastname,
		&user.Password,
		&user.CreatedAt,
	)
	if err != nil {
		return &user, err
	}
	return &user, nil
}

func (u *UserStorage) Save(ctx context.Context, user *UserRegister) (*User, error) {
	var createdUser User
	query := `
										INSERT INTO users (username, lastname, password, created_at)
										VALUES ($1, $2, $3, NOW())
										RETURNING id, username, lastname, created_at;`
	generatedPasswordBytes, _ := bcrypt.GenerateFromPassword([]byte(user.Password), 10)
	generatedPassword := string(generatedPasswordBytes)
	fmt.Println(generatedPassword)

	row := u.db.Pool.QueryRow(ctx, query, &user.Username, &user.Lastname, generatedPassword)
	err := row.Scan(&createdUser.ID, &createdUser.Username, &createdUser.Lastname, &createdUser.CreatedAt)
	if err != nil {
		return &createdUser, err
	}
	return &createdUser, nil
}
