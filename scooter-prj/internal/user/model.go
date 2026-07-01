package user

import "time"

type User struct {
	ID        int       `db:"id" json:"id"`
	Username  string    `db:"username" json:"username"`
	Lastname  string    `db:"lastname" json:"lastname"`
	Password  string    `db:"password" json:"password"`
	CreatedAt time.Time `db:"created-at" json:"created-at"`
}
