package handlers

import (
	"context"
	"fmt"
	"trainmate/dao"
	"trainmate/models"
)

type JobHandler interface {
	InsertOne(models.Job) interface{}
	FindOne(string) models.Job
	FindByName(string) models.Job
	FindByNameExpid(string, string) models.Job
	FindAll() []models.Job
	DeleteOne(string) error
}

// ***************************** mongo *********************
type MongoJobHandler struct {
	TblName string
}

// "Job"
func NewMongoJobHandler(tblName string) *MongoJobHandler {
	return &MongoJobHandler{TblName: tblName}
}

// 插入单个
func (m *MongoJobHandler) InsertOne(obj models.Job) (insertResult interface{}) {
	insertResult = dao.NewMgo(m.TblName).InsertOne(obj)
	// if err != nil {
	// 	fmt.Println(err)
	// }
	return
}

func (m *MongoJobHandler) FindOne(id string) models.Job {
	var result models.Job
	dao.NewMgo(m.TblName).FindOne("Id", id).Decode(&result)
	return result
}

func (m *MongoJobHandler) FindByName(name string) models.Job {
	var result models.Job
	dao.NewMgo(m.TblName).FindOne("name", name).Decode(&result)
	return result
}

func (m *MongoJobHandler) FindByNameExpid(name string, expid string) models.Job {
	// var result models.Job
	// dao.NewMgo(m.TblName).FindOne("name", name).Decode(&result)
	// return result
	var result models.Job
	var filterMap = make(map[string]interface{})
	filterMap["name"] = name
	filterMap["expid"] = expid
	dao.NewMgo(m.TblName).FindOneByD(filterMap).Decode((&result))
	// dao.NewMgo(jobTblName).FindOne("name", m.Name).Decode(&result)
	return result
}

func (m *MongoJobHandler) FindAll() []models.Job {
	var results []models.Job
	var filterMap = make(map[string]interface{})
	cursor := dao.NewMgo(m.TblName).FindMany(filterMap)
	for cursor.Next(context.Background()) {
		ud := &models.Job{}
		err := cursor.Decode(ud)
		results = append(results, *ud)
		if err != nil {
			fmt.Println("decode error is ", err)
			continue
		}
	}
	return results
}

func (m *MongoJobHandler) DeleteOne(id string) error {
	_, err := dao.NewMgo(m.TblName).Delete("id", id)
	return err
}
