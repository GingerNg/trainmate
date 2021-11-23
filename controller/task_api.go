package controller

import (
	"fmt"
	"trainmate/initialize"
	"trainmate/models/request"
	"trainmate/models/response"

	"github.com/gin-gonic/gin"
)

func CreateTask(c *gin.Context) {

	var params request.TaskParams
	err := c.ShouldBindJSON(&params) // bind

	if err != nil {
		response.WriteHttpErrMsg(c, 999, "job保存失败! bind error", err.Error())
		return
	}
	obj, err := initialize.TaskSvc.CreateTask(&params)

	if err != nil {
		response.WriteHttpErrMsg(c, 999, "job保存失败!", err.Error())
		return
	}
	// logger.Infof("文件上传耗时: %s", time.Since(start))
	// utils.WriteHttpOkMsgs(c, id, "")
	response.WriteHttpOkMsgs(c, obj.Id, "")
}

// func UpdateTask(c *gin.Context) {

// 	var params request.TaskQueryParams
// 	err := c.ShouldBindJSON(&params) // bind

// 	if err != nil {
// 		response.WriteHttpErrMsg(c, 999, "job更新失败!", err.Error())
// 		return
// 	}
// 	job, err := Service.UpdateTask(&params)

// 	if err != nil {
// 		response.WriteHttpErrMsg(c, 999, "job更新失败!", err.Error())
// 		return
// 	}
// 	// logger.Infof("文件上传耗时: %s", time.Since(start))
// 	// utils.WriteHttpOkMsgs(c, id, "")
// 	response.WriteHttpOkMsgs(c, job.Id, job)
// }

func QueryTask(c *gin.Context) {
	var params request.TaskQueryParams
	err := c.ShouldBind(&params) // bind

	fmt.Println(params)

	if err != nil {
		response.WriteHttpErrMsg(c, 999, "task查询失败!", err.Error())
		return
	}
	job, err := initialize.TaskSvc.QueryTask(&params)

	if err != nil {
		response.WriteHttpErrMsg(c, 999, "task query error!", err.Error())
		return
	}
	// logger.Infof("文件上传耗时: %s", time.Since(start))
	// utils.WriteHttpOkMsgs(c, id, "")
	response.WriteHttpOkMsgs(c, job.Id, job)
}

func QueryTasks(c *gin.Context) {
	var params request.TaskQueryParams
	err := c.ShouldBind(&params) // bind

	fmt.Println(params)

	if err != nil {
		response.WriteHttpErrMsg(c, 999, "tasks query 失败!", err.Error())
		return
	}
	objs := initialize.TaskSvc.QueryTasks(&params)

	if err != nil {
		response.WriteHttpErrMsg(c, 999, "tasks query 失败!", err.Error())
		return
	}
	// logger.Infof("文件上传耗时: %s", time.Since(start))
	// utils.WriteHttpOkMsgs(c, id, "")
	response.WriteHttpOkMsgs(c, "", objs)
}

func DeleteTask(c *gin.Context) {
	var params request.TaskQueryParams
	err := c.ShouldBindJSON(&params) // bind

	if err != nil {
		response.WriteHttpErrMsg(c, 999, "task 删除失败!", err.Error())
		return
	}
	err = initialize.TaskSvc.DeleteTask(&params)

	if err != nil {
		response.WriteHttpErrMsg(c, 999, "task 删除失败!", err.Error())
		return
	}
	// logger.Infof("文件上传耗时: %s", time.Since(start))
	// utils.WriteHttpOkMsgs(c, id, "")
	response.WriteHttpOkMsgs(c, "", "")
}
