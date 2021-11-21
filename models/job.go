package models

type Job struct {
	Id         string      `json:"id"`
	Name       string      `json:"name"`
	ExpId      string      `json:"exp_id"`
	Metrics    interface{} `json:"metrics"`
	ModelUrl   string      `json:"model_url"`
	Config     interface{} `json:"config"`
	History    interface{} `json:"history"`
	Updatetime string      `json:"update_time"`
	Createtime string      `json:"create_time"`
	Status     int         `json:"status"`
}

// var jobTblName = "Job"
