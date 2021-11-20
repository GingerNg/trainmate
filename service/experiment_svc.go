package service

import (
	"fmt"
	"trainmate/models"
	"trainmate/models/request"
	"trainmate/utils"
)

func CreateExp(params *request.ExperimentParams) (models.Experiment, error) {
	exp := models.Experiment{
		Id:        utils.GetUUID(),
		Name:      params.Name,
		DatasetId: params.DatasetId}
	res := exp.FindByName()
	if res.Id == "" {
		exp.InsertOne()
		return exp, nil
	} else {
		fmt.Println("duplicate") // warning
		return res, nil
	}
}

func QueryExp(params *request.ExperimentQueryParams) (models.Experiment, error) {
	obj := models.Experiment{Id: params.Id}
	result := obj.FindOne()
	return result, nil
}

func QueryExps(params *request.ExperimentQueryParams) []models.Experiment {
	exp := models.Experiment{}
	results := exp.FindAll()
	return results
}

func DeleteExp(params *request.ExperimentQueryParams) error {
	obj := models.Experiment{Id: params.Id}
	err := obj.DeleteOne()
	return err
}
