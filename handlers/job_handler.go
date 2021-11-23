package handlers

import (
	"context"
	"fmt"
	"trainmate/dao"
	"trainmate/models"

	"gorm.io/gorm"
)

type JobHandler interface {
	InsertOne(models.Job) interface{}
	FindOne(string) models.Job
	FindByName(string) models.Job
	FindByNameExpid(string, string) models.Job
	FindAll() []models.Job
	DeleteOne(string) error
}

// ***************************** sqlite *******************
type RdbJobHandler struct {
	// config *config.Config
	db *gorm.DB
}

func NewRdbJobHandler(db *gorm.DB) *RdbJobHandler {
	return &RdbJobHandler{db: db}
}

func (r *RdbJobHandler) InsertOne(obj models.Job) (insertResult interface{}) {
	// insertResult = dao.NewMgo(m.TblName).InsertOne(obj)
	// if err != nil {
	// 	fmt.Println(err)
	// }
	result := r.db.Create(&obj)
	if err := result.Error; err != nil {
		fmt.Println(err)
		return obj
	}
	return obj
}

func (r *RdbJobHandler) FindOne(id string) models.Job {
	var obj models.Job
	if err := r.db.Where(&models.Job{Id: id}).First(&obj).Error; err != nil {
		return obj
	}
	return obj
}

func (r *RdbJobHandler) FindByNameExpid(name string, id string) models.Job {
	var obj models.Job
	if err := r.db.Where(&models.Job{Name: name, Id: id}).First(&obj).Error; err != nil {
		fmt.Println(err)
		return obj
	}
	return obj
}

func (r *RdbJobHandler) FindByName(name string) models.Job {
	var obj models.Job
	if err := r.db.Where(&models.Job{Name: name}).First(&obj).Error; err != nil {
		fmt.Println(err)
		return obj
	}
	return obj
}

func (r *RdbJobHandler) FindAll() []models.Job {
	var objs []models.Job
	if err := r.db.Find(&objs).Error; err != nil {
		return objs
	}
	return objs
}

func (r *RdbJobHandler) DeleteOne(id string) error {
	// _, err := dao.NewMgo(m.TblName).Delete("id", id)
	return r.db.
		Where("id = ?", id).
		Delete(models.Job{}).Error
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
