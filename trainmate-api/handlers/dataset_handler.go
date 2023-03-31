package handlers

import (
	"context"
	"fmt"
	"trainmate/dao"
	"trainmate/models"

	"gorm.io/gorm"
)

type DatasetHandler interface {
	InsertOne(models.Dataset) interface{}
	FindOne(string) models.Dataset
	FindByName(string) models.Dataset
	FindAll() []models.Dataset
	DeleteOne(string) error
}

// **************************** sqlite ********************
type RdbDatasetHandler struct {
	// config *config.Config
	db *gorm.DB
}

func NewRdbDatasetHandler(db *gorm.DB) *RdbDatasetHandler {
	return &RdbDatasetHandler{db: db}
}

func (r *RdbDatasetHandler) InsertOne(obj models.Dataset) (insertResult interface{}) {
	// insertResult = dao.NewMgo(m.TblName).InsertOne(obj)
	// if err != nil {
	// 	fmt.Println(err)
	// }
	result := r.db.Create(&obj)
	if err := result.Error; err != nil {
		return obj
	}
	return obj
}

func (r *RdbDatasetHandler) FindOne(id string) models.Dataset {
	var obj models.Dataset
	if err := r.db.Where(&models.Dataset{Id: id}).First(&obj).Error; err != nil {
		return obj
	}
	return obj
}

func (r *RdbDatasetHandler) FindByName(name string) models.Dataset {
	var obj models.Dataset
	if err := r.db.Where(&models.Dataset{Name: name}).First(&obj).Error; err != nil {
		fmt.Println(err)
		return obj
	}
	return obj
}

func (r *RdbDatasetHandler) FindAll() []models.Dataset {
	var objs []models.Dataset
	if err := r.db.Find(&objs).Error; err != nil {
		return objs
	}
	return objs
}

func (r *RdbDatasetHandler) DeleteOne(id string) error {
	// _, err := dao.NewMgo(m.TblName).Delete("id", id)
	return r.db.
		Where("id = ?", id).
		Delete(models.Dataset{}).Error
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
