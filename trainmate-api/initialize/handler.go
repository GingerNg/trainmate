package initialize

import (
	"trainmate/handlers"
)

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

	} else {
		TaskHdl = handlers.NewRdbTaskHandler(db)
		DatasetHdl = handlers.NewRdbDatasetHandler(db)
		ExpHdl = handlers.NewRdbExperimentHandler(db)
		JobHdl = handlers.NewRdbJobHandler(db)
	}

	// Init Handler

}
