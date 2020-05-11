package users

import (
	"encoding/json"
)

type PublicUser struct {
	ID             int64     `json:"id"`
	DateCreated    string    `json:"date_created"`
	Status         string    `json:"status"`
}

type PrivateUser struct {
	ID             int64     `json:"id"`
	FirstName      string    `json:"first_name"`
	LastName       string    `json:"last_name"`
	Email          string    `json:"email"`
	DateCreated    string    `json:"date_created"`
	Status         string    `json:"status"`
}

func (users Users) Marshall(isPublic bool) []interface{} {
	result := make([]interface{}, len(users))
	for i, user := range users {
		result[i] = user.Marshall(isPublic)
	}
	return result
}

func (u *User) Marshall(isPublic bool) interface{} {
	if isPublic {
		return PublicUser{
			ID:           u.ID,
			DateCreated:  u.DateCreated,
			Status:       u.Status,
		}
	}
	// Mapping JSON struct User to struct of type PrivateUser
	userJson, _ := json.Marshal(u)
	var privateUser PrivateUser
	_ = json.Unmarshal(userJson, &privateUser)
	return privateUser
}
