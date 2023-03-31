package models

type Account struct {
	Id         string
	Name       string
	Createtime CustomTime `gorm:"type:timestamp; default:CURRENT_TIMESTAMP" swaggertype:"string" format:"date" example:"2006-01-02 15:04:05.000"`
}

type User struct {
	// Account Account
	ID             string     `json:"id" gorm:"primary_key"`
	ApiKey         string     `json:"api_key" gorm:"unique"`
	Email          string     `json:"email" gorm:"index:idx_user_email; size:255"`
	Location       string     `json:"location"`
	Password       string     `json:"-"`
	LastLoggedInAt CustomTime `gorm:"type:timestamp; default:CURRENT_TIMESTAMP" swaggertype:"string"`
}

type Group struct {
	Account   Account
	UserCount int64
}

type UserGroupRelation struct {
	UserId  string
	GroupId string
}

type ExperimentRelation struct {
	Id        string
	ExpId     string
	AccountId string
	RelType   string // 关联关系类型 1. 属于，2. 分享
}
