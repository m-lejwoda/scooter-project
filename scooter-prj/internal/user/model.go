package user

import "time"

type User struct {
	ID        int       `db:"id"`
	Username  string    `db:"username"`
	Lastname  string    `db:"lastname"`
	Password  string    `db:"password"`
	CreatedAt time.Time `db:"created-at"`
}
