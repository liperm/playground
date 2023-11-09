package model

import "github.com/google/uuid"

type User struct {
	Id          uuid.UUID
	Name        string
	Email       string
	AccessLevel string
}
