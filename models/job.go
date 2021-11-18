package models

import (
	"context"
	"fmt"
	"trainmate/dao"
)

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

var jobTblName = "Job"

// 插入单个
func (m *Job) InsertOne() (insertResult interface{}) {
	insertResult = dao.NewMgo(jobTblName).InsertOne(m)
	// if err != nil {
	// 	fmt.Println(err)
	// }
	return
}

func (m *Job) UpdateOne() (updateResult interface{}) {

	// todo
	update := make(map[string]interface{})
	if m.Config != "" {
		update["config"] = m.Config
	}
	if m.History != "" {
		update["history"] = m.History
	}
	if m.Metrics != "" {
		update["metrics"] = m.Metrics
	}
	if m.ModelUrl != "" {
		update["modelurl"] = m.ModelUrl
	}

	updateResult = dao.NewMgo(jobTblName).UpdateOne("id", m.Id, update)
	return
}

func (m *Job) FindOne() Job {
	var result Job
	dao.NewMgo(jobTblName).FindOne("id", m.Id).Decode(&result)
	return result
}

func (m *Job) FindByName() Job {
	var result Job
	var filterMap = make(map[string]interface{})
	filterMap["name"] = m.Name
	filterMap["expid"] = m.ExpId
	dao.NewMgo(jobTblName).FindOneByD(filterMap).Decode((&result))
	// dao.NewMgo(jobTblName).FindOne("name", m.Name).Decode(&result)
	return result
}

func (m *Job) FindByExpid() []Job {
	var filterMap = make(map[string]interface{})
	if m.ExpId != "" {
		filterMap["expid"] = m.ExpId
	}
	cursor := dao.NewMgo(jobTblName).FindMany(filterMap)
	var results []Job
	for cursor.Next(context.Background()) {
		ud := &Job{}
		err := cursor.Decode(ud)
		results = append(results, *ud)
		if err != nil {
			fmt.Println("decode error is ", err)
			continue
		}
	}
	return results
}

func (m *Job) DeleteOne() error {
	_, err := dao.NewMgo(jobTblName).Delete("id", m.Id)
	return err
}
