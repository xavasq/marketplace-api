package domain

type User struct {
	ID       int64
	Email    string
	Password string
	Role     string
	Balance  int64
}
