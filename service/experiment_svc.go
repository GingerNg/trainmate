package service

import (
	"fmt"
	"trainmate/handlers"
	"trainmate/models"
	"trainmate/models/request"
	"trainmate/utils"
)

type ExpService struct {
	handler handlers.ExperimentHandler
}

var ExpSvc ExpService

func NewExpService() *ExpService {
	return &ExpService{handler: handlers.NewMongoExperimentHandler("Experiment")}
}

func init() {
	ExpSvc = *NewExpService()
}

func (m *ExpService) CreateExp(params *request.ExperimentParams) (models.Experiment, error) {
	obj := models.Experiment{
		Id:        utils.GetUUID(),
		Name:      params.Name,
		DatasetId: params.DatasetId}
	res := m.handler.FindByName(obj.Name)
	if res.Id == "" {
		m.handler.InsertOne(obj)
		return obj, nil
	} else {
		fmt.Println("duplicate") // warning
		return res, nil
	}
}

func (m *ExpService) QueryExp(params *request.ExperimentQueryParams) (models.Experiment, error) {
	result := m.handler.FindOne(params.Id)
	return result, nil
}

func (m *ExpService) QueryExps(params *request.ExperimentQueryParams) []models.Experiment {
	results := m.handler.FindAll()
	return results
}

func (m *ExpService) DeleteExp(params *request.ExperimentQueryParams) error {
	err := m.handler.DeleteOne(params.Id)
	return err
}
