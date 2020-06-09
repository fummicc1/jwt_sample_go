package model

// User is an Entity that stores about user.
type User struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Posts []Post `json:"posts"`
}
