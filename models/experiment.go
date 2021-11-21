package models

type Experiment struct {
	Id        string `json:"id"` // job-id
	Name      string `json:"name"`
	DatasetId string `json:"dataset_id"`
	Desp      string `json:"desp"`
	Status    int    `json:"status"`
}

// var exprimentTblName = "Experiment"
