package database

import (
	"errors"
	"time"

	"github.com/google/uuid"
)

// Post -
type Post struct {
	ID        string    `json:"id"`
	CreatedAt time.Time `json:"createdAt"`
	UserEmail string    `json:"userEmail"`
	Text      string    `json:"text"`
}

// CreatePost -
func (c Client) CreatePost(userEmail, text string) (Post, error) {
	db, err := c.readDatabase()
	if err != nil {
		return Post{}, err
	}
	_, ok := db.Users[userEmail]
	if !ok {
		return Post{}, errors.New("user doesn't exist")
	}

	post := Post{
		ID:        uuid.New().String(),
		CreatedAt: time.Now().UTC(),
		UserEmail: userEmail,
		Text:      text,
	}
	db.Posts[post.ID] = post
	err = c.updateDatabase(db)
	if err != nil {
		return Post{}, err
	}
	return post, nil
}

// GetPosts -
func (c Client) GetPosts(userEmail string) ([]Post, error) {
	db, err := c.readDatabase()
	if err != nil {
		return nil, err
	}
	_, ok := db.Users[userEmail]
	if !ok {
		return nil, errors.New("user doesn't exist")
	}
	var posts []Post
	for _, post := range db.Posts {
		if post.UserEmail == userEmail {
			posts = append(posts, post)
		}
	}
	return posts, nil
}

// DeletePost -
func (c Client) DeletePost(id string) error {
	db, err := c.readDatabase()
	if err != nil {
		return err
	}
	_, ok := db.Posts[id]
	if !ok {
		return errors.New("post doesn't exist")
	}
	delete(db.Posts, id)
	err = c.updateDatabase(db)
	if err != nil {
		return err
	}
	return nil
}
