package initialize

import "trainmate/service"

var (
	TaskSvc    *service.TaskService
	DatasetSvc *service.DatasetService
	ExpSvc     *service.ExpService
	JobSvc     *service.JobService
)

func InitService() {
	TaskSvc = service.NewTaskService(TaskHdl)
	DatasetSvc = service.NewDatasetService(DatasetHdl)
	ExpSvc = service.NewExpService(ExpHdl)
	JobSvc = service.NewJobServicee(JobHdl)
}
