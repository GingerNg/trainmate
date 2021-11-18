package request

type ExperimentQueryParams struct {
	Id string `json:"exp_id" form:"exp_id"` // job-id
}

type ExperimentParams struct {
	Id        string `json:"exp_id" form:"exp_id"` // exp-id
	DatasetId string `json:"dataset_id" form:"dataset_id"`
	Name      string `json:"name" form:"name"`
}
