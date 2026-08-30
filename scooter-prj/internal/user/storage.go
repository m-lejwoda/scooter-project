// Repository implementation
package user

import (
	"context"
	"errors"
	"fmt"
	"time"

	"scooter-prj/internal/database"

	"github.com/jackc/pgx/v5"
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

func (u *UserStorage) GetByEmail(ctx context.Context, email string) (*User, error) {
	query := `SELECT id, username, lastname, password, created_at FROM users Where email = $1`
	var user User
	err := u.db.Pool.QueryRow(ctx, query, email).Scan(
		&user.ID,
		&user.Username,
		&user.Lastname,
		&user.Password,
		&user.CreatedAt,
	)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, ErrUserNotFound
		}
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

func (u *UserStorage) SavePasswordResetToken(ctx context.Context, userId int, tokenHash string) (*PasswordResetTokenResponse, error) {
	query := `INSERT INTO password_reset_tokens (user_id, token_hash, expired_at, created_at)
									VALUES ($1, $2, $3, $4, NOW())
									RETURNING id, user_id, token_hash, expired_at, created_at`
	expiredAt := time.Now().Add(time.Minute * 30)
	row := u.db.Pool.QueryRow(ctx, query, userId, tokenHash, expiredAt)
	var resetToken PasswordResetTokenResponse
	err := row.Scan(&resetToken.ID, &resetToken.UserId, &resetToken.TokenHash, &resetToken.ExpiredAt, &resetToken.CreatedAt)
	if err != nil {
		fmt.Println("Error")
		return nil, err
	}
	return &resetToken, nil
}
