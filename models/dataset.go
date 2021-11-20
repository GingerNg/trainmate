package models

import (
	"context"
	"fmt"
	"trainmate/dao"
)

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

var datasetTblName = "Dataset"

// 插入单个
func (m *Dataset) InsertOne() (insertResult interface{}) {
	insertResult = dao.NewMgo(datasetTblName).InsertOne(m)
	// if err != nil {
	// 	fmt.Println(err)
	// }
	return
}

func (m *Dataset) FindByName() Dataset {
	var result Dataset
	dao.NewMgo(datasetTblName).FindOne("name", m.Name).Decode(&result)
	return result
}

func (m *Dataset) FindOne() Dataset {
	var result Dataset
	dao.NewMgo(datasetTblName).FindOne("Id", m.Id).Decode(&result)
	return result
}

func (m *Dataset) FindAll() []Dataset {
	var results []Dataset
	var filterMap = make(map[string]interface{})
	cursor := dao.NewMgo(datasetTblName).FindMany(filterMap)
	for cursor.Next(context.Background()) {
		ud := &Dataset{}
		err := cursor.Decode(ud)
		results = append(results, *ud)
		if err != nil {
			fmt.Println("decode error is ", err)
			continue
		}
	}
	return results
}

func (m *Dataset) DeleteOne() error {
	_, err := dao.NewMgo(datasetTblName).Delete("id", m.Id)
	return err
}
