package service

import (
	"errors"
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
		return exp, errors.New("duplicate")
	}
}

func QueryExp(params *request.ExperimentQueryParams) []models.Experiment {
	exp := models.Experiment{}
	results := exp.FindAll()
	return results
}
