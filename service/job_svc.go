package service

import (
	"fmt"
	"trainmate/handlers"
	"trainmate/models"
	"trainmate/models/request"
	"trainmate/utils"
)

type JobService struct {
	handler handlers.JobHandler
}

var JobSvc JobService

func NewJobServicee() *JobService {
	return &JobService{handler: handlers.NewMongoJobHandler("Job")}
}

func init() {
	JobSvc = *NewJobServicee()
}

func (m *JobService) CreateJob(params *request.JobParams) (models.Job, error) {
	obj := models.Job{Id: utils.GetUUID(),
		Name:    params.Name,
		ExpId:   params.ExpId,
		Metrics: params.Metrics,
		Config:  params.Config,
		History: params.History,
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

func (m *JobService) UpdateJob(params *request.JobParams) (models.Job, error) {
	obj := models.Job{
		Id:       params.Id,
		Name:     params.Name,
		Metrics:  params.Metrics,
		History:  params.History,
		ModelUrl: params.ModelUrl,
		Config:   params.Config,
	}
	// m.handler.DeleteOne().UpdateOne()
	return obj, nil
}

func (m *JobService) QueryJob(params *request.JobQueryParams) (models.Job, error) {
	if params.ExpId != "" && params.Name != "" {
		res := m.handler.FindByNameExpid(params.Name, params.ExpId)
		return res, nil
	} else {
		res := m.handler.FindOne(params.Id)
		return res, nil
	}
}

func (m *JobService) QueryJobs(params *request.JobQueryParams) []models.Job {
	results := m.handler.FindAll()
	return results
}

func (m *JobService) DeleteJob(params *request.JobQueryParams) error {
	err := m.handler.DeleteOne(params.Id)
	return err
}
