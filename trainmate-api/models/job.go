package models

type Job struct {
	Id         string     `json:"id"`
	Name       string     `json:"name"`
	ExpId      string     `json:"exp_id"`
	Metrics    string     `json:"metrics"`
	ModelUrl   string     `json:"model_url"`
	Config     string     `json:"config"`
	History    string     `json:"history"`
	Updatetime CustomTime `json:"update_time" gorm:"type:timestamp; default:CURRENT_TIMESTAMP" swaggertype:"string" format:"date" example:"2006-01-02 15:04:05.000"`
	Createtime CustomTime `json:"create_time" gorm:"type:timestamp; default:CURRENT_TIMESTAMP" swaggertype:"string" format:"date" example:"2006-01-02 15:04:05.000"`
	Status     int        `json:"status"`
}

// var jobTblName = "Job"
