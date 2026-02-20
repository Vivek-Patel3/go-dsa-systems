package auth

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

type User struct {
	userName       string
	hashedPassword []byte
}

type UserLoginSystem struct {
	userStore map[string]*User
}

func NewUser(userName, password string) *User {
	hashedPassword,_ := bcrypt.GenerateFromPassword([]byte(password), 10)
	
	return &User{
		userName:       userName,
		hashedPassword: hashedPassword,
	}
}

func NewUserLoginSystem() *UserLoginSystem {
	return &UserLoginSystem{
		userStore: make(map[string]*User),
	}
}

func (loginSystem *UserLoginSystem) AddUser(userName, password string) {
	user := NewUser(userName, password)
	loginSystem.userStore[user.userName] = user
}

func (loginSystem *UserLoginSystem) LoginUser(userName, password string) (*User, bool) {
	if val, ok := loginSystem.userStore[userName]; ok {
		if err := bcrypt.CompareHashAndPassword(val.hashedPassword, []byte(password)); err == nil {
			return val, true
		} else {
			fmt.Println("Incorrect Password")
			return nil, false
		}
	} else {
		fmt.Println("No such user exists")
		return nil, false 
	}
}