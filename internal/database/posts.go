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
	data, err := c.readDatabase()
	if err != nil {
		return Post{}, err
	}
	_, ok := data.Users[userEmail]
	if !ok {
		return Post{}, errors.New("user doesn't exist")
	}

	post := Post{
		ID:        uuid.New().String(),
		CreatedAt: time.Now().UTC(),
		UserEmail: userEmail,
		Text:      text,
	}
	data.Posts[post.ID] = post
	err = c.updateDatabase(data)
	if err != nil {
		return Post{}, err
	}
	return post, nil
}

// GetPosts -
func (c Client) GetPosts(userEmail string) ([]Post, error) {
	data, err := c.readDatabase()
	if err != nil {
		return nil, err
	}
	_, ok := data.Users[userEmail]
	if !ok {
		return nil, errors.New("user doesn't exist")
	}
	var posts []Post
	for _, post := range data.Posts {
		if post.UserEmail == userEmail {
			posts = append(posts, post)
		}
	}
	return posts, nil
}

// DeletePost -
func (c Client) DeletePost(id string) error {
	data, err := c.readDatabase()
	if err != nil {
		return err
	}
	_, ok := data.Posts[id]
	if !ok {
		return errors.New("post doesn't exist")
	}
	delete(data.Posts, id)
	err = c.updateDatabase(data)
	if err != nil {
		return err
	}
	return nil
}
