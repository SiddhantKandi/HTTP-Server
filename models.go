package main

import (
	"time"

	"github.com/SiddhantKandi/HTTPServer/internal/database"
	"github.com/google/uuid"
)

type User struct {
	ID        uuid.UUID  `json:"id"`
	CreatedAt time.Time  `json:"createdAt"`
	UpdatedAt time.Time  `json:"updatedAt"`
	Name      string     `json:"name"`
}

func databaseUsertoUser(dbUser database.User) User{
	return User{
		ID:dbUser.ID,
		CreatedAt:dbUser.CreatedAt,
		UpdatedAt:dbUser.UpdatedAt,
		Name:dbUser.Name,
	}
}