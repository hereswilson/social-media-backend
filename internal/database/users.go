package database

import (
	"errors"
	"time"
)

// User -
type User struct {
	CreatedAt time.Time `json:"createdAt"`
	Email     string    `json:"email"`
	Password  string    `json:"password"`
	Name      string    `json:"name"`
	Age       int       `json:"age"`
}

// CreateUser -
func (c Client) CreateUser(email, password, name string, age int) (User, error) {
	data, err := c.readDatabase()
	if err != nil {
		return User{}, err
	}
	user := User{
		CreatedAt: time.Now().UTC(),
		Email:     email,
		Password:  password,
		Name:      name,
		Age:       age,
	}
	data.Users[email] = user
	err = c.updateDatabase(data)
	if err != nil {
		return User{}, err
	}
	return user, nil
}

func (c Client) UpdateUser(email, password, name string, age int) (User, error) {
	data, err := c.readDatabase()
	if err != nil {
		return User{}, err
	}
	user, ok := data.Users[email]
	if !ok {
		return User{}, errors.New("user doesn't exist")
	}

	user.Password = password
	user.Name = name
	user.Age = age

	data.Users[email] = user

	err = c.updateDatabase(data)
	if err != nil {
		return User{}, err
	}
	return user, nil
}

func (c Client) GetUser(email string) (User, error) {
	data, err := c.readDatabase()
	if err != nil {
		return User{}, err
	}
	user, ok := data.Users[email]
	if !ok {
		return User{}, errors.New("user doesn't exist")
	}
	return user, nil
}

func (c Client) DeleteUser(email string) error {
	data, err := c.readDatabase()
	if err != nil {
		return err
	}
	_, ok := data.Users[email]
	if !ok {
		return errors.New("user doesn't exist")
	}
	delete(data.Users, email)
	err = c.updateDatabase(data)
	if err != nil {
		return err
	}
	return nil
}
