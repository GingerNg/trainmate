package handlers

import (
	"context"
	"fmt"
	"trainmate/dao"
)

type Handler interface {
	InsertOne(interface{}) interface{}
	FindOne(string) interface{}
	FindByName(string) interface{}
	FindAll() interface{}
	DeleteOne(string) error
}

// ***************************** mongo *********************
type MongoHandler struct {
	TblName string
}

func NewMongoHandler(tblName string) *MongoHandler {
	return &MongoHandler{TblName: tblName}
}

// 插入单个
func (m *MongoHandler) InsertOne(obj interface{}) (insertResult interface{}) {
	insertResult = dao.NewMgo(m.TblName).InsertOne(obj)
	// if err != nil {
	// 	fmt.Println(err)
	// }
	return
}

func (m *MongoHandler) FindOne(id string) interface{} {
	var result interface{}
	dao.NewMgo(m.TblName).FindOne("Id", id).Decode(&result)
	return result
}

func (m *MongoHandler) FindByName(name string) interface{} {
	var obj interface{}
	dao.NewMgo(m.TblName).FindOne("name", name).Decode(&obj)
	return obj
}

func (m *MongoHandler) FindAll(results interface{}) interface{} {

	var filterMap = make(map[string]interface{})
	cursor := dao.NewMgo(m.TblName).FindMany(filterMap)
	if err := cursor.All(context.TODO(), &results); err != nil {
		fmt.Println(err)
	}
	// fmt.Println("---")
	return results
}

func (m *MongoHandler) DeleteOne(id string) error {
	_, err := dao.NewMgo(m.TblName).Delete("id", id)
	return err
}
