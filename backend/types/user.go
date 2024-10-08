package types

import (
	"time"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

// User model
type User struct {
	ID                uuid.UUID `json:"id"`
	FirstName         string    `json:"first_name"`
	LastName          string    `json:"last_name"`
	UserName          string    `json:"user_name"`
	EncryptedPassword string    `json:"encrypted_password"`
	CreatedAt         time.Time `json:"created_at"`
	Workouts          []string  `json:"workouts"`
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
	}, nil
}

type UserPayload struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	UserName  string `json:"user_name"`
	Password  string `json:"password"`
}

func (p *UserPayload) Validate() map[string]string {
	errs := make(map[string]string)
	return errs
}
