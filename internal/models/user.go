package models

type User struct {
	ID       string `db:"id"`
	Username string `db:"username"`
	Emaill   string `db:"email"`
	Password string `db:"password"`
}