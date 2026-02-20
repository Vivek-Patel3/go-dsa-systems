package main

import "github.com/VivekPatel3/go-dsa-systems/auth"

type ChatAppBackend struct {
	users map[string]*auth.User
	loginService *auth.UserLoginSystem
}

func NewChatApp() *ChatAppBackend {
	return &ChatAppBackend{
		users: make(map[string]*auth.User),
		loginService: auth.NewUserLoginSystem(),
	}
}

func (app *ChatAppBackend) RegisterUser(userName, password string) {
	app.loginService.AddUser(userName, password)
}

