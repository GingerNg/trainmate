package request

type TaskQueryParams struct {
	Id   string `json:"task_id" form:"task_id"` // exp-id
	Name string `json:"name"  form:"name"`
}

type TaskParams struct {
	Id   string `json:"task_id" form:"task_id"` // exp-id
	Name string `json:"name" form:"name"`
}
