package handlers

import "trainmate/models"

type ExperimentHandler interface {
	GetAll() ([]*models.Dataset, error)
	GetById(uint) (*models.Dataset, error)
	GetByUser(string) ([]*models.Dataset, error)
	Insert(*models.Dataset) (*models.Dataset, error)
	Delete(uint) error
}
