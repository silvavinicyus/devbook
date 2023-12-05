package models

import (
	"errors"
	"strings"
	"time"
)

type Post struct {
	ID          uint64    `json:"id,omitempty"`
	Title       string    `json:"title,omitempty"`
	Content     string    `json:"content,omitempty"`
	CreatorId   uint64    `json:"creatorId,omitempty"`
	CreatorNick string    `json:"creatorNick,omitempty"`
	CreatedAt   time.Time `json:"createdAt,omitempty"`
	Likes       uint64    `json:"likes"`
}

func (post *Post) Prepare() error {
	if erro := post.validate(); erro != nil {
		return erro
	}

	post.format()

	return nil
}

func (post *Post) validate() error {
	if post.Title == "" {
		return errors.New("title shoudn't be empty")
	}

	if post.Content == "" {
		return errors.New("title shoudn't be empty")
	}

	return nil
}

func (post *Post) format() {
	post.Title = strings.TrimSpace(post.Title)
	post.Content = strings.TrimSpace(post.Content)
}
