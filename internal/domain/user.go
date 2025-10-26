package domain

import "time"

type User struct {
	ID       string
	Name     string
	Email    string
	Password string
	CreatedAt time.Time
}

type UserRepository interface {
	Create(user *User) error
	FindByEmail(email string) (*User, error)
}