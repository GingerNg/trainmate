package initialize

import "trainmate/service"

var (
	TaskSvc    *service.TaskService
	DatasetSvc *service.DatasetService
	ExpSvc     *service.ExpService
	JobSvc     *service.JobService
)

func InitService() {
	// Init Handler
	TaskSvc = service.NewTaskService()
	DatasetSvc = service.NewDatasetService()
	ExpSvc = service.NewExpService()
	JobSvc = service.NewJobServicee()
}
