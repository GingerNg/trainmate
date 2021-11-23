package handlers

import "trainmate/models"

type UserHandler interface {
	InsertOne(models.User) interface{}
	FindOne(string) models.User
	// FindByName(string) models.Dataset
	// FindAll() []models.Dataset
	DeleteOne(string) error
}
