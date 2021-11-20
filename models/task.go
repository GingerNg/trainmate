package models

import (
	"context"
	"fmt"
	"trainmate/dao"
)

type Task struct {
	Id       string `json:"id"`
	Name     string `json:"name"`
	Version  string `json:"version"`
	Desp     string `json:"desp"`
	MetaInfo string `json:"meta_info"`
}

var taskTblName = "Task"

// 插入单个
func (m *Task) InsertOne() (insertResult interface{}) {
	insertResult = dao.NewMgo(taskTblName).InsertOne(m)
	// if err != nil {
	// 	fmt.Println(err)
	// }
	return
}

func (m *Task) FindOne() Task {
	var result Task
	dao.NewMgo(taskTblName).FindOne("Id", m.Id).Decode(&result)
	return result
}

func (m *Task) FindByName() Task {
	var result Task
	dao.NewMgo(taskTblName).FindOne("name", m.Name).Decode(&result)
	return result
}

func (m *Task) FindAll() []Task {
	var results []Task
	var filterMap = make(map[string]interface{})
	cursor := dao.NewMgo(taskTblName).FindMany(filterMap)
	for cursor.Next(context.Background()) {
		ud := &Task{}
		err := cursor.Decode(ud)
		results = append(results, *ud)
		if err != nil {
			fmt.Println("decode error is ", err)
			continue
		}
	}
	return results
}

func (m *Task) DeleteOne() error {
	_, err := dao.NewMgo(taskTblName).Delete("id", m.Id)
	return err
}
