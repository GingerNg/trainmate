package service

import (
	"errors"
	"trainmate/models"
	"trainmate/models/request"
	"trainmate/utils"
)

func CreateJob(params *request.JobParams) (models.Job, error) {
	job := models.Job{Id: utils.GetUUID(),
		Name:    params.Name,
		ExpId:   params.ExpId,
		Metrics: params.Metrics,
		Config:  params.Config,
		History: params.History,
	}
	res := job.FindByName()
	if res.Id == "" {
		job.InsertOne()
		return job, nil
	} else {
		return job, errors.New("duplicate")
	}
}

func UpdateJob(params *request.JobParams) (models.Job, error) {
	job := models.Job{
		Id:       params.Id,
		Name:     params.Name,
		Metrics:  params.Metrics,
		History:  params.History,
		ModelUrl: params.ModelUrl,
		Config:   params.Config,
	}
	job.UpdateOne()
	return job, nil
}

func QueryJob(params *request.JobQueryParams) (models.Job, error) {
	var job models.Job
	if params.ExpId != "" && params.Name != "" {
		job = models.Job{Name: params.Name, ExpId: params.ExpId}
		res := job.FindByName()
		return res, nil
	} else {
		job = models.Job{Id: params.Id}
		res := job.FindOne()
		return res, nil
	}
}

func QueryJobByExpid(params *request.JobQueryParams) []models.Job {
	job := models.Job{ExpId: params.ExpId}
	results := job.FindByExpid()
	return results
}

func DeleteJob(params *request.JobQueryParams) error {
	job := models.Job{Id: params.Id}
	err := job.DeleteOne()
	return err
}
