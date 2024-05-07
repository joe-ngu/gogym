package user

import (
	"time"

	"golang.org/x/crypto/bcrypt"
)

// User model
type User struct {
	ID                uint
	FirstName         string
	LastName          string
	UserName          string
	EncryptedPassword string
	CreatedAt         time.Time
	Workouts          []uint
}

func (u *User) ValidPassword(pw string) bool {
	return bcrypt.CompareHashAndPassword([]byte(u.EncryptedPassword), []byte(pw)) == nil
}

func NewUser(firstName string, lastName string, userName string, password string) (*User, error) {
	encpw, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	return &User{
		FirstName:         firstName,
		LastName:          lastName,
		UserName:          userName,
		EncryptedPassword: string(encpw),
		CreatedAt:         time.Now().UTC(),
	}, nil
}

type LoginResponse struct {
	UserName string `json:"user_name"`
	Token    string `json:"token"`
}

type LoginRequest struct {
	UserName string `json:"user_name"`
	Password string `json:"password"`
}

type CreateAccountRequest struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Password  string `json:"password"`
}


