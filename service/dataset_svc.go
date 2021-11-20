package service

import (
	"fmt"
	"trainmate/models"
	"trainmate/models/request"
	"trainmate/utils"
)

func CreateDataset(params *request.DatasetParams) (models.Dataset, error) {
	obj := models.Dataset{
		Id:     utils.GetUUID(),
		Name:   params.Name,
		Desp:   params.Desp,
		TaskId: params.TaskId,
	}
	res := obj.FindByName()
	if res.Id == "" {
		obj.InsertOne()
		return obj, nil
	} else {
		fmt.Println("duplicate") // warning
		return res, nil
	}
}

func QueryDataset(params *request.DatasetQueryParams) (models.Dataset, error) {
	obj := models.Dataset{Id: params.Id}
	result := obj.FindOne()
	return result, nil
}

func QueryDatasets(params *request.DatasetQueryParams) []models.Dataset {
	obj := models.Dataset{}
	results := obj.FindAll()
	return results
}

func DeleteDataset(params *request.DatasetQueryParams) error {
	obj := models.Dataset{Id: params.Id}
	err := obj.DeleteOne()
	return err
}
