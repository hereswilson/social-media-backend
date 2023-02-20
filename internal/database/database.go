package database

import (
	"encoding/json"
	"errors"
	"os"
	"time"
)

type Client struct {
	dbPath string
}

func NewClient(dbPath string) Client {
	return Client{dbPath: dbPath}
}

type databaseSchema struct {
	Users map[string]User `json:"users"`
	Posts map[string]Post `json:"posts"`
}

// User -
type User struct {
	CreatedAt time.Time `json:"createdAt"`
	Email     string    `json:"email"`
	Password  string    `json:"password"`
	Name      string    `json:"name"`
	Age       int       `json:"age"`
}

// Post -
type Post struct {
	ID        string    `json:"id"`
	CreatedAt time.Time `json:"createdAt"`
	UserEmail string    `json:"userEmail"`
	Text      string    `json:"text"`
}

// createDatabase -
func (c Client) createDatabase() error {
	data, err := json.Marshal(databaseSchema{
		Users: make(map[string]User),
		Posts: make(map[string]Post),
	})
	if err != nil {
		return err
	}
	err = os.WriteFile(c.dbPath, data, 0600)
	if err != nil {
		return err
	}
	return nil
}

// EnsureDatabase -
func (c Client) EnsureDatabase() error {
	_, err := os.Stat(c.dbPath)
	if os.IsNotExist(err) {
		return c.createDatabase()
	}
	return err
}

func (c Client) updateDatabase(data databaseSchema) error {
	jsonData, err := json.Marshal(data)
	if err != nil {
		return err
	}
	err = os.WriteFile(c.dbPath, jsonData, 0600)
	if err != nil {
		return err
	}
	return nil
}

func (c Client) readDatabase() (databaseSchema, error) {
	data := databaseSchema{}
	file, err := os.Open(c.dbPath)
	if err != nil {
		return data, err
	}
	defer file.Close()
	err = json.NewDecoder(file).Decode(&data)
	if err != nil {
		return data, err
	}
	return data, nil
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
