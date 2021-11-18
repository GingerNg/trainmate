package models

import (
	"trainmate/dao"
)

type Dataset struct {
	Id       string
	TaskId   string
	Name     string
	Version  string
	Desp     string
	Source   string
	Url      string
	MetaInfo string
}

var datasetTblName = "Dataset"

// 插入单个
func (m *Dataset) InsertOne() (insertResult interface{}) {
	insertResult = dao.NewMgo(datasetTblName).InsertOne(m)
	// if err != nil {
	// 	fmt.Println(err)
	// }
	return
}

func (m *Dataset) FindOne() Dataset {
	var result Dataset
	dao.NewMgo(datasetTblName).FindOne("Id", m.Id).Decode(&result)
	return result
}
