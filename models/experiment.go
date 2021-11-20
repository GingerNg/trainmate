package models

import (
	"context"
	"fmt"
	"trainmate/dao"
)

type Experiment struct {
	Id        string `json:"id"` // job-id
	Name      string `json:"name"`
	DatasetId string `json:"dataset_id"`
	Desp      string `json:"desp"`
	Status    int    `json:"status"`
}

var exprimentTblName = "Experiment"

// 插入单个
func (m *Experiment) InsertOne() (insertResult interface{}) {
	insertResult = dao.NewMgo(exprimentTblName).InsertOne(m)
	// if err != nil {
	// 	fmt.Println(err)
	// }
	return
}

func (m *Experiment) FindOne() Experiment {
	var result Experiment
	dao.NewMgo(exprimentTblName).FindOne("id", m.Id).Decode(&result)
	return result
}

func (m *Experiment) FindByName() Experiment {
	var result Experiment
	var filterMap = make(map[string]interface{})
	filterMap["name"] = m.Name
	dao.NewMgo(exprimentTblName).FindOneByD(filterMap).Decode((&result))
	// dao.NewMgo(jobTblName).FindOne("name", m.Name).Decode(&result)
	return result
}

func (m *Experiment) FindAll() []Experiment {
	var filterMap = make(map[string]interface{})
	// filterMap["expid"] = m.ExpId
	cursor := dao.NewMgo(exprimentTblName).FindMany(filterMap)
	var results []Experiment
	for cursor.Next(context.Background()) {
		ud := &Experiment{}
		err := cursor.Decode(ud)
		results = append(results, *ud)
		if err != nil {
			fmt.Println("decode error is ", err)
			continue
		}
	}
	return results
}

func (m *Experiment) DeleteOne() error {
	_, err := dao.NewMgo(exprimentTblName).Delete("id", m.Id)
	return err
}
