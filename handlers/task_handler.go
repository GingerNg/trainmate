package handlers

import (
	"context"
	"fmt"
	"trainmate/dao"
	"trainmate/models"
)

type TaskHandler interface {
	InsertOne(models.Task) interface{}
	FindOne(string) models.Task
	FindByName(string) models.Task
	FindAll() []models.Task
	DeleteOne(string) error
}

// ***************************** mongo *********************
type MongoTaskHandler struct {
	TblName string
}

func NewMongoTaskHandler() *MongoTaskHandler {
	return &MongoTaskHandler{TblName: "Task"}
}

// 插入单个
func (m *MongoTaskHandler) InsertOne(obj models.Task) (insertResult interface{}) {
	insertResult = dao.NewMgo(m.TblName).InsertOne(obj)
	// if err != nil {
	// 	fmt.Println(err)
	// }
	return
}

func (m *MongoTaskHandler) FindOne(id string) models.Task {
	var result models.Task
	dao.NewMgo(m.TblName).FindOne("Id", id).Decode(&result)
	return result
}

func (m *MongoTaskHandler) FindByName(name string) models.Task {
	var result models.Task
	dao.NewMgo(m.TblName).FindOne("name", name).Decode(&result)
	return result
}

func (m *MongoTaskHandler) FindAll() []models.Task {
	var results []models.Task
	var filterMap = make(map[string]interface{})
	cursor := dao.NewMgo(m.TblName).FindMany(filterMap)
	for cursor.Next(context.Background()) {
		ud := &models.Task{}
		err := cursor.Decode(ud)
		results = append(results, *ud)
		if err != nil {
			fmt.Println("decode error is ", err)
			continue
		}
	}
	return results
}

func (m *MongoTaskHandler) DeleteOne(id string) error {
	_, err := dao.NewMgo(m.TblName).Delete("id", id)
	return err
}
