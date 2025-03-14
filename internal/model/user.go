package model

import (
	"fmt"
	"github.com/vlbarou/sampleproject/internal/serializer"
)

type User struct {
	ID       uint `gorm:"primaryKey;autoIncrement"`
	Name     string
	Username string
	Email    string
}

func NewUser(id uint, name, username, password string) *User {

	return &User{
		ID:       id,
		Name:     name,
		Username: username,
		Email:    password,
	}
}

func (u *User) setName(name string) {
	u.Name = name
}

func (u *User) setUsername(username string) {
	u.Username = username
}

func (u *User) setEmail(email string) {
	u.Email = email
}

func (u *User) getName() string {
	return u.Name
}

func (u *User) getUsername() string {
	return u.Username
}

func (u *User) getEmail() string {
	return u.Email
}

// Implement the Stringer interface: fmt.Println(p) Calls p.String() automatically
func (u User) String() string {
	return fmt.Sprintf("User: " + serializer.MarshalStructOrEmpty(u))
}
