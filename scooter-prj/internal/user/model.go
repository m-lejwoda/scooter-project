package user

import "time"

type User struct {
	ID        int       `db:"id" json:"id"`
	Username  string    `db:"username" json:"username"`
	Lastname  string    `db:"lastname" json:"lastname"`
	Password  string    `db:"password" json:"password"`
	CreatedAt time.Time `db:"created_at" json:"created_at"`
}

type PasswordResetToken struct {
	ID        int       `db:"id" json:"id"`
	UserID    int       `db:"user_id" json:"user_id"`
	TokenHash string    `db:"token_hash" json:"-"`
	ExpiredAt time.Time `db:"expired_at" json:"expired_at"`
	CreatedAt time.Time `db:"created_at" json:"created_at"`
}
