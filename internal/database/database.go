package database

import (
	"encoding/json"
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
