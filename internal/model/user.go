package model

import "fmt"

type User struct {
	Name     string
	Username string
	Email    string
}

func NewUser(name, username, password string) *User {

	return &User{
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

// Implement the Stringer interface: fmt.Println(p) // Calls p.String() automatically
func (u User) String() string {
	return fmt.Sprintf("User{Name: %s, Username: %d}", u.Name, u.Username)
}
