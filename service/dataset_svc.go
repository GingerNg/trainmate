package service

import (
	"fmt"
	"trainmate/handlers"
	"trainmate/models"
	"trainmate/models/request"
	"trainmate/utils"
)

type DatasetService struct {
	handler handlers.DatasetHandler
}

var DatasetSvc DatasetService

func NewDatasetService() *DatasetService {
	return &DatasetService{handler: handlers.NewMongoDatasetHandler("Dataset")}
}

func init() {
	DatasetSvc = *NewDatasetService()
}

func (m *DatasetService) CreateDataset(params *request.DatasetParams) (models.Dataset, error) {
	obj := models.Dataset{
		Id:     utils.GetUUID(),
		Name:   params.Name,
		Desp:   params.Desp,
		TaskId: params.TaskId,
	}
	res := m.handler.FindByName(obj.Name)
	if res.Id == "" {
		m.handler.InsertOne(obj)
		return obj, nil
	} else {
		fmt.Println("duplicate") // warning
		return res, nil
	}
}

func (m *DatasetService) QueryDataset(params *request.DatasetQueryParams) (models.Dataset, error) {
	result := m.handler.FindOne(params.Id)
	return result, nil
}

func (m *DatasetService) QueryDatasets(params *request.DatasetQueryParams) []models.Dataset {
	results := m.handler.FindAll()
	return results
}

func (m *DatasetService) DeleteDataset(params *request.DatasetQueryParams) error {
	err := m.handler.DeleteOne(params.Id)
	return err
}
