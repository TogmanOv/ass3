package domain

import "time"

type User struct {
	id           int64
	Details      UserDetails
	Email        string
	PasswordHash string
	createdAt    time.Time
}

func (u *User) ID() int64 {
	return u.id
}

func (u *User) CreatedAt() time.Time {
	return u.createdAt
}

type UserDetails struct {
	FirstName string
	LastName  string
	Phone     string
}
