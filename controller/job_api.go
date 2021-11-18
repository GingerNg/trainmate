package controller

import (
	"fmt"
	"trainmate/models/request"
	"trainmate/models/response"
	Service "trainmate/service"

	"github.com/gin-gonic/gin"
)

func CreateJob(c *gin.Context) {

	var params request.JobParams
	err := c.ShouldBindJSON(&params) // bind

	if err != nil {
		response.WriteHttpErrMsg(c, 999, "job保存失败! bind error", err.Error())
		return
	}
	job, err := Service.CreateJob(&params)

	if err != nil {
		response.WriteHttpErrMsg(c, 999, "job保存失败!", err.Error())
		return
	}
	// logger.Infof("文件上传耗时: %s", time.Since(start))
	// utils.WriteHttpOkMsgs(c, id, "")
	response.WriteHttpOkMsgs(c, job.Id, "")
}

func UpdateJob(c *gin.Context) {

	var params request.JobParams
	err := c.ShouldBindJSON(&params) // bind

	if err != nil {
		response.WriteHttpErrMsg(c, 999, "job更新失败!", err.Error())
		return
	}
	job, err := Service.UpdateJob(&params)

	if err != nil {
		response.WriteHttpErrMsg(c, 999, "job更新失败!", err.Error())
		return
	}
	// logger.Infof("文件上传耗时: %s", time.Since(start))
	// utils.WriteHttpOkMsgs(c, id, "")
	response.WriteHttpOkMsgs(c, job.Id, job)
}

func QueryJob(c *gin.Context) {
	var params request.JobQueryParams
	err := c.ShouldBind(&params) // bind

	fmt.Println(params)

	if err != nil {
		response.WriteHttpErrMsg(c, 999, "job查询失败!", err.Error())
		return
	}
	job, err := Service.QueryJob(&params)

	if err != nil {
		response.WriteHttpErrMsg(c, 999, "job保存失败!", err.Error())
		return
	}
	// logger.Infof("文件上传耗时: %s", time.Since(start))
	// utils.WriteHttpOkMsgs(c, id, "")
	response.WriteHttpOkMsgs(c, job.Id, job)
}

func QueryJobByExp(c *gin.Context) {
	var params request.JobQueryParams
	err := c.ShouldBind(&params) // bind

	if err != nil {
		response.WriteHttpErrMsg(c, 999, "job查询失败, bind error!", err.Error())
		return
	}
	jobs := Service.QueryJobByExpid(&params)

	if err != nil {
		response.WriteHttpErrMsg(c, 999, "job查询失败, query error!", err.Error())
		return
	}
	// logger.Infof("文件上传耗时: %s", time.Since(start))
	// utils.WriteHttpOkMsgs(c, id, "")
	response.WriteHttpOkMsgs(c, "", jobs)
}

func DeleteJob(c *gin.Context) {
	var params request.JobQueryParams
	err := c.ShouldBindJSON(&params) // bind

	if err != nil {
		response.WriteHttpErrMsg(c, 999, "job删除失败!", err.Error())
		return
	}
	err = Service.DeleteJob(&params)

	if err != nil {
		response.WriteHttpErrMsg(c, 999, "job删除失败!", err.Error())
		return
	}
	// logger.Infof("文件上传耗时: %s", time.Since(start))
	// utils.WriteHttpOkMsgs(c, id, "")
	response.WriteHttpOkMsgs(c, "", "")
}
