package service

import (
	"fmt"
	"trainmate/handlers"
	"trainmate/models"
	"trainmate/models/request"
	"trainmate/utils"
)

type TaskService struct {
	handler handlers.TaskHandler
}

var TaskSvc TaskService

func NewTaskService(taskHandler handlers.TaskHandler) *TaskService {
	// return &TaskService{handler: handlers.NewMongoTaskHandler("Task")}
	return &TaskService{handler: taskHandler}
}

// func init() {
// 	TaskSvc = *NewTaskService()
// }

func (m *TaskService) CreateTask(params *request.TaskParams) (models.Task, error) {
	obj := models.Task{
		Id:   utils.GetUUID(),
		Name: params.Name}
	res := m.handler.FindByName(obj.Name)
	if res.Id == "" {
		m.handler.InsertOne(obj)
		return obj, nil
	} else {
		fmt.Println("duplicate") // warning
		return res, nil
	}
}

func (m *TaskService) QueryTask(params *request.TaskQueryParams) (models.Task, error) {
	result := m.handler.FindOne(params.Id)
	return result, nil
}

func (m *TaskService) QueryTasks(params *request.TaskQueryParams) []models.Task {
	results := m.handler.FindAll()
	return results
}

func (m *TaskService) DeleteTask(params *request.TaskQueryParams) error {
	err := m.handler.DeleteOne(params.Id)
	return err
}
