package request

type DatasetQueryParams struct {
	Id   string `json:"dataset_id" form:"dataset_id"` // job-id
	Name string `json:"name"  form:"name"`
}

type DatasetParams struct {
	Id     string `json:"dataset_id" form:"dataset_id"` // exp-id
	TaskId string `json:"task_id" form:"task_id"`
	Name   string `json:"name" form:"name"`
	Desp   string `json:"desp" form:"desp"`
}
