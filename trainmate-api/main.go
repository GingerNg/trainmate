package main

// import (
// 	"context"
// 	"fmt"
// 	"log"
// 	"trainmate/conf"
// 	"trainmate/dao"
// 	"trainmate/drivers"
// 	"trainmate/handlers"
// 	"trainmate/models"

// 	"github.com/gin-gonic/gin"
// 	"go.mongodb.org/mongo-driver/bson"
// )

// // var Engine *gin.Engine

// func RegisterRouter1() {
// 	// Engine.POST("/upload/", handlers.Upload)
// 	// Engine.GET("/query/", handlers.Query)
// 	// Engine.GET("/download/", handlers.Download)
// 	Engine.GET("/hello/", handlers.Hello)

// 	// distributedstroage
// 	PerfRouter()

// 	// Engine.POST("/distributed/storage", handlers.Upload)
// }

// func PerfRouter() {
// 	group := Engine.Group("/trainmate/perf")
// 	{
// 		// group.POST("insert/", handlers.DfsUpload)
// 		group.GET("query/", handlers.Hello)
// 		// group.POST("/callback/notice/", controllers.JobsCtr.CallbackNotice)
// 		// group.POST("/re_parsed/", controllers.JobsCtr.ReParsed)
// 	}
// }

// type Trainer struct {
// 	Name string
// 	Age  int
// 	City string
// }

// func main1() {

// 	// 初始化数据库
// 	drivers.Init()

// 	perf := models.Perf{Id: "12121", JobId: "1212", Metrics: "232"}
// 	perf.InsertOne()

// 	obj := perf.FindOneById()
// 	fmt.Printf("Found a single document: %+v\n", obj)

// 	// 单个插入
// 	ash := Trainer{"Ash", 10, "Pallet Town"}
// 	//  mgo_dao := dao.NewMgo()
// 	InsertOneResult := dao.NewMgo("trainer").InsertOne(ash)
// 	fmt.Println("Inserted a single document: ", InsertOneResult)

// 	// 插入多个值
// 	misty := Trainer{"Misty", 10, "Cerulean City"}
// 	brock := Trainer{"Brock", 15, "Pewter City"}
// 	trainers := []interface{}{misty, brock}
// 	insertManyResult := dao.NewMgo("trainer").InsertMany(trainers)
// 	fmt.Println("Inserted multiple documents: ", insertManyResult)

// 	// 更新
// 	update := bson.D{
// 		{"$inc", bson.D{
// 			{"age", 999},
// 		}},
// 	}
// 	updateResult := dao.NewMgo("trainer").UpdateOne("name", "Ash", update)
// 	fmt.Printf("Matched %v documents and updated %v documents.\n", updateResult.MatchedCount, updateResult.ModifiedCount)

// 	// 查询一个
// 	var result Trainer
// 	dao.NewMgo("trainer").FindOne("name", "Ash").Decode(&result)
// 	fmt.Printf("Found a single document: %+v\n", result)

// 	// 查询总数
// 	name, size := dao.NewMgo("trainer").Count()
// 	fmt.Printf(" documents name: %+v documents size %d \n", name, size)

// 	// 查询多个记录
// 	var results []*Trainer
// 	cur := dao.NewMgo("trainer").FindAll(0, size, 1)
// 	defer cur.Close(context.TODO())
// 	if cur != nil {
// 		fmt.Println("FindAll err:", cur)
// 	}
// 	for cur.Next(context.TODO()) {
// 		var elem Trainer
// 		err := cur.Decode(&elem)
// 		if err != nil {
// 			log.Fatal(err)
// 		}
// 		results = append(results, &elem)
// 	}
// 	if err := cur.Err(); err != nil {
// 		log.Fatal(err)
// 	}
// 	// 遍历结果
// 	for k, v := range results {

// 		fmt.Printf("Found  documents  %d  %v \n", k, v)
// 	}

// 	// 删除文件
// 	deleteResult := dao.NewMgo("trainer").DeleteMany("name", "Ash")
// 	fmt.Printf("Deleted %v documents in the trainers collection\n", deleteResult)

// 	// 关闭连接
// 	// drivers.Close()

// 	Engine = gin.New()
// 	//  Engine.Use(middleware.CostStatMiddleware())
// 	// 注册orm映射
// 	//  models.InitOrm()
// 	// 注册路由信息
// 	// RegisterRouter()
// 	err := Engine.Run(conf.GetGlobalConfig().GetTcpBind())
// 	if err != nil {
// 		panic(err)
// 	}

// }
