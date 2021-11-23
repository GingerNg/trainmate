package request

type JobQueryParams struct {
	Id    string `json:"job_id"  form:"job_id"` // job-id
	ExpId string `json:"exp_id"  form:"exp_id"`
	Name  string `json:"name"  form:"name"`
}

type JobParams struct {
	Id       string `json:"job_id" form:"job_id"` // job-id
	ExpId    string `json:"exp_id" form:"exp_id"`
	Name     string `json:"name"  form:"name" binding:"required"`
	Metrics  string `json:"metrics" form:"metrics"`
	ModelUrl string `json:"model_url" form:"model_url"`
	Config   string `json:"config"  form:"config"`
	History  string `json:"history" form:"history"`
}
