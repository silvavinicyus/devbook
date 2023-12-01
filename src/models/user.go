package models

import (
	"api/src/security"
	"errors"
	"strings"
	"time"

	"github.com/badoux/checkmail"
)

type User struct {
	ID        uint64    `json:"id,omitempty"`
	Name      string    `json:"name,omitempty"`
	Nick      string    `json:"nick,omitempty"`
	Email     string    `json:"email,omitempty"`
	Password  string    `json:"Password,omitempty"`
	CreatedAt time.Time `json:"created_at,omitempty"`
}

func (u *User) Prepare(step string) error {
	if erro := u.validate(step); erro != nil {
		return erro
	}

	erro := u.format(step)
	if erro != nil {
		return erro
	}

	return nil
}

func (u *User) validate(step string) error {
	if u.Name == "" {
		return errors.New("name is a mandatory parameter and should not be empty")
	}

	if u.Nick == "" {
		return errors.New("nick is a mandatory parameter and should not be empty")
	}

	if u.Email == "" {
		return errors.New("email is a mandatory parameter and should not be empty")
	}

	if erro := checkmail.ValidateFormat(u.Email); erro != nil {
		return errors.New("this email is not valid")
	}

	if step == "signup" && u.Password == "" {
		return errors.New("password is a mandatory parameter and should not be empty")
	}

	return nil
}

func (u *User) format(step string) error {
	u.Name = strings.TrimSpace(u.Name)
	u.Email = strings.TrimSpace(u.Email)
	u.Nick = strings.TrimSpace(u.Nick)

	if step == "signup" {
		hashedPassword, erro := security.Hash(u.Password)

		if erro != nil {
			return erro
		}

		u.Password = string(hashedPassword)
	}

	return nil
}
