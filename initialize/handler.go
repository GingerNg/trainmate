package initialize

import "trainmate/handlers"

var (
	TaskHdl    handlers.TaskHandler
	DatasetHdl handlers.DatasetHandler
	ExpHdl     handlers.ExperimentHandler
	JobHdl     handlers.JobHandler
)

func InitHandler() {
	if config.UseMongo {
		TaskHdl = handlers.NewMongoTaskHandler("Task")
		DatasetHdl = handlers.NewMongoDatasetHandler("Dataset")
		ExpHdl = handlers.NewMongoExperimentHandler("Experiment")
		JobHdl = handlers.NewMongoJobHandler("Job")

	} else if config.Db.Type == "sqlite" {
		TaskHdl = handlers.NewRdbTaskHandler(db)
	}
	// Init Handler

}
