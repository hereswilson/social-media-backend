package database

import (
	"encoding/json"
	"os"
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
