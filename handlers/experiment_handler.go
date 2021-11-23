package handlers

import (
	"context"
	"fmt"
	"trainmate/dao"
	"trainmate/models"

	"gorm.io/gorm"
)

type ExperimentHandler interface {
	InsertOne(models.Experiment) interface{}
	FindOne(string) models.Experiment
	FindByName(string) models.Experiment
	FindAll() []models.Experiment
	DeleteOne(string) error
}

// ***************************** sqlite *******************
type RdbExperimentHandler struct {
	// config *config.Config
	db *gorm.DB
}

func NewRdbExperimentHandler(db *gorm.DB) *RdbExperimentHandler {
	return &RdbExperimentHandler{db: db}
}

func (r *RdbExperimentHandler) InsertOne(obj models.Experiment) (insertResult interface{}) {
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

func (r *RdbExperimentHandler) FindOne(id string) models.Experiment {
	var obj models.Experiment
	if err := r.db.Where(&models.Experiment{Id: id}).First(&obj).Error; err != nil {
		return obj
	}
	return obj
}

func (r *RdbExperimentHandler) FindByName(name string) models.Experiment {
	var obj models.Experiment
	if err := r.db.Where(&models.Experiment{Name: name}).First(&obj).Error; err != nil {
		fmt.Println(err)
		return obj
	}
	return obj
}

func (r *RdbExperimentHandler) FindAll() []models.Experiment {
	var objs []models.Experiment
	if err := r.db.Find(&objs).Error; err != nil {
		return objs
	}
	return objs
}

func (r *RdbExperimentHandler) DeleteOne(id string) error {
	// _, err := dao.NewMgo(m.TblName).Delete("id", id)
	return r.db.
		Where("id = ?", id).
		Delete(models.Experiment{}).Error
}

// ***************************** mongo *********************
type MongoExperimentHandler struct {
	TblName string
}

func NewMongoExperimentHandler(tblName string) *MongoExperimentHandler {
	return &MongoExperimentHandler{TblName: tblName}
}

// 插入单个
func (m *MongoExperimentHandler) InsertOne(obj models.Experiment) (insertResult interface{}) {
	insertResult = dao.NewMgo(m.TblName).InsertOne(obj)
	// if err != nil {
	// 	fmt.Println(err)
	// }
	return
}

func (m *MongoExperimentHandler) FindOne(id string) models.Experiment {
	var result models.Experiment
	dao.NewMgo(m.TblName).FindOne("Id", id).Decode(&result)
	return result
}

func (m *MongoExperimentHandler) FindByName(name string) models.Experiment {
	var result models.Experiment
	dao.NewMgo(m.TblName).FindOne("name", name).Decode(&result)
	return result
}

func (m *MongoExperimentHandler) FindAll() []models.Experiment {
	var results []models.Experiment
	var filterMap = make(map[string]interface{})
	cursor := dao.NewMgo(m.TblName).FindMany(filterMap)
	for cursor.Next(context.Background()) {
		ud := &models.Experiment{}
		err := cursor.Decode(ud)
		results = append(results, *ud)
		if err != nil {
			fmt.Println("decode error is ", err)
			continue
		}
	}
	return results
}

func (m *MongoExperimentHandler) DeleteOne(id string) error {
	_, err := dao.NewMgo(m.TblName).Delete("id", id)
	return err
}
