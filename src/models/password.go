package models

type PasswordUpdate struct {
	New     string `json:"new"`
	Current string `json:"current"`
}
