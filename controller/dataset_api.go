package controller

import (
	"fmt"
	"trainmate/models/request"
	"trainmate/models/response"
	Service "trainmate/service"

	"github.com/gin-gonic/gin"
)

func CreateDataset(c *gin.Context) {

	var params request.DatasetParams
	err := c.ShouldBindJSON(&params) // bind

	if err != nil {
		response.WriteHttpErrMsg(c, 999, "job保存失败! bind error", err.Error())
		return
	}
	job, err := Service.CreateDataset(&params)

	if err != nil {
		response.WriteHttpErrMsg(c, 999, "job保存失败!", err.Error())
		return
	}
	// logger.Infof("文件上传耗时: %s", time.Since(start))
	// utils.WriteHttpOkMsgs(c, id, "")
	response.WriteHttpOkMsgs(c, job.Id, "")
}

func UpdateDataset(c *gin.Context) {

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

func QueryDataset(c *gin.Context) {
	var params request.DatasetQueryParams
	err := c.ShouldBind(&params) // bind

	fmt.Println(params)

	if err != nil {
		response.WriteHttpErrMsg(c, 999, "job查询失败!", err.Error())
		return
	}
	job, err := Service.QueryDataset(&params)

	if err != nil {
		response.WriteHttpErrMsg(c, 999, "job保存失败!", err.Error())
		return
	}
	// logger.Infof("文件上传耗时: %s", time.Since(start))
	// utils.WriteHttpOkMsgs(c, id, "")
	response.WriteHttpOkMsgs(c, job.Id, job)
}

func QueryDatasets(c *gin.Context) {
	var params request.DatasetQueryParams
	err := c.ShouldBind(&params) // bind

	if err != nil {
		response.WriteHttpErrMsg(c, 999, "job查询失败, bind error!", err.Error())
		return
	}
	jobs := Service.QueryDatasets(&params)

	if err != nil {
		response.WriteHttpErrMsg(c, 999, "job查询失败, query error!", err.Error())
		return
	}
	// logger.Infof("文件上传耗时: %s", time.Since(start))
	// utils.WriteHttpOkMsgs(c, id, "")
	response.WriteHttpOkMsgs(c, "", jobs)
}

func DeleteDataset(c *gin.Context) {
	var params request.DatasetQueryParams
	err := c.ShouldBindJSON(&params) // bind

	if err != nil {
		response.WriteHttpErrMsg(c, 999, "job删除失败!", err.Error())
		return
	}
	err = Service.DeleteDataset(&params)

	if err != nil {
		response.WriteHttpErrMsg(c, 999, "job删除失败!", err.Error())
		return
	}
	// logger.Infof("文件上传耗时: %s", time.Since(start))
	// utils.WriteHttpOkMsgs(c, id, "")
	response.WriteHttpOkMsgs(c, "", "")
}
