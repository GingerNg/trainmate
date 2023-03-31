package models

import (
	"trainmate/dao"
)

type Perf struct {
	Id string
	JobId  string
	Metrics string
	// Dao *dao.Mgo
}

var perfTblName = "Perf"


 // 插入单个
 func (m *Perf) InsertOne() (insertResult interface{}) {
	insertResult = dao.NewMgo(perfTblName).InsertOne(m)
	// if err != nil {
	// 	fmt.Println(err)
	// }
	return
}

func (m *Perf) FindOneById()(Perf){
	var result Perf
	dao.NewMgo(perfTblName).FindOne("Id", m.Id).Decode(&result)
	return result
}


