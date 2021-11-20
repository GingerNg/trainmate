package handlers

import (
	"context"
	"fmt"
	"trainmate/dao"
	"trainmate/models"
)

type DatasetHandler interface {
	InsertOne(models.Dataset) interface{}
	FindOne(string) models.Dataset
	FindByName(string) models.Dataset
	FindAll() []models.Dataset
	DeleteOne(string) error
}

// ***************************** mongo *********************
type MongoDatasetHandler struct {
	TblName string
}

// "Dataset"
func NewMongoDatasetHandler(tblName string) *MongoDatasetHandler {
	return &MongoDatasetHandler{TblName: tblName}
}

// 插入单个
func (m *MongoDatasetHandler) InsertOne(obj models.Dataset) (insertResult interface{}) {
	insertResult = dao.NewMgo(m.TblName).InsertOne(obj)
	// if err != nil {
	// 	fmt.Println(err)
	// }
	return
}

func (m *MongoDatasetHandler) FindOne(id string) models.Dataset {
	var result models.Dataset
	dao.NewMgo(m.TblName).FindOne("Id", id).Decode(&result)
	return result
}

func (m *MongoDatasetHandler) FindByName(name string) models.Dataset {
	var result models.Dataset
	dao.NewMgo(m.TblName).FindOne("name", name).Decode(&result)
	return result
}

func (m *MongoDatasetHandler) FindAll() []models.Dataset {
	var results []models.Dataset
	var filterMap = make(map[string]interface{})
	cursor := dao.NewMgo(m.TblName).FindMany(filterMap)
	for cursor.Next(context.Background()) {
		ud := &models.Dataset{}
		err := cursor.Decode(ud)
		results = append(results, *ud)
		if err != nil {
			fmt.Println("decode error is ", err)
			continue
		}
	}
	return results
}

func (m *MongoDatasetHandler) DeleteOne(id string) error {
	_, err := dao.NewMgo(m.TblName).Delete("id", id)
	return err
}
