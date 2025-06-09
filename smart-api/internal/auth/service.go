package auth

import (
	"encoding/json"
	"errors"
	"os"
)

const userFile = "./data/user.json"

func LoadUser() ([]User, error) {
	file, err := os.Open(userFile)

	if err != nil {
		return nil, err
	}
	defer file.Close()

	var users []User
	err = json.NewDecoder(file).Decode(&users)
	if err != nil {
		return nil, err
	}

	return users, nil
}

func SaveUsers(users []User) error {
	file, err := os.Create(userFile)
	if err != nil {
		return err
	}
	defer file.Close()

	return json.NewEncoder(file).Encode(users)
}

func RegisterUser(username, password string) error {
	users, err := LoadUser()
	if err != nil {
		return err
	}

	for _, user := range users {
		if user.Username == username {
			return errors.New("user already exists")
		}
	}

	newUser := User{
		ID: len(users) + 1,
		Username: username,
		Password: password,
	}
	
	users = append(users, newUser)
	return SaveUsers(users)
}

func LoginUser(username, password string) (*User, error) {
	users, err := LoadUser()
	if err != nil {
		return nil, err
	}

	for _, user  := range users {
		if user.Username == username && user.Password == password {
			return &user, nil
		}
	}

	return nil, errors.New("invalid credentials")
}