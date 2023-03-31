package handlers

import (
	"context"
	"fmt"
	"trainmate/dao"
	"trainmate/models"

	// "github.com/revel/config"
	"gorm.io/gorm"
)

// "Task"
type TaskHandler interface {
	InsertOne(models.Task) interface{} // 插入单个
	FindOne(string) models.Task
	FindByName(string) models.Task
	FindAll() []models.Task
	DeleteOne(string) error
}

// ***************************** sqlite *******************
type RdbTaskHandler struct {
	// config *config.Config
	db *gorm.DB
}

func NewRdbTaskHandler(db *gorm.DB) *RdbTaskHandler {
	return &RdbTaskHandler{db: db}
}

func (r *RdbTaskHandler) InsertOne(obj models.Task) (insertResult interface{}) {
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

func (r *RdbTaskHandler) FindOne(id string) models.Task {
	var obj models.Task
	if err := r.db.Where(&models.Task{Id: id}).First(&obj).Error; err != nil {
		return obj
	}
	return obj
}

func (r *RdbTaskHandler) FindByName(name string) models.Task {
	var obj models.Task
	if err := r.db.Where(&models.Task{Name: name}).First(&obj).Error; err != nil {
		fmt.Println(err)
		return obj
	}
	return obj
}

func (r *RdbTaskHandler) FindAll() []models.Task {
	var objs []models.Task
	if err := r.db.Find(&objs).Error; err != nil {
		return objs
	}
	return objs
}

func (r *RdbTaskHandler) DeleteOne(id string) error {
	// _, err := dao.NewMgo(m.TblName).Delete("id", id)
	return r.db.
		Where("id = ?", id).
		Delete(models.Task{}).Error
}

// ***************************** mongo *********************
type MongoTaskHandler struct {
	TblName string
}

func NewMongoTaskHandler(tblName string) *MongoTaskHandler {
	return &MongoTaskHandler{TblName: tblName}
}

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
