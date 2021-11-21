package models

type Dataset struct {
	Id       string `json:"id"`
	TaskId   string `json:"task_id"`
	Name     string `json:"name"`
	Version  string `json:"version"`
	Desp     string `json:"desp"`
	Source   string `json:"source"`
	Url      string `json:"url"`
	MetaInfo string `json:"meta_info"`
}

// var datasetTblName = "Dataset"
