package controller

import (
	"fmt"
	"trainmate/initialize"
	"trainmate/models/request"
	"trainmate/models/response"

	"github.com/gin-gonic/gin"
)

func CreateExp(c *gin.Context) {

	var params request.ExperimentParams
	err := c.ShouldBindJSON(&params) // bind

	if err != nil {
		response.WriteHttpErrMsg(c, 999, "exp保存失败!", err.Error())
		return
	}
	exp, err := initialize.ExpSvc.CreateExp(&params)

	if err != nil {
		response.WriteHttpErrMsg(c, 999, "exp保存失败!", err.Error())
		return
	}
	// logger.Infof("文件上传耗时: %s", time.Since(start))
	// utils.WriteHttpOkMsgs(c, id, "")
	response.WriteHttpOkMsgs(c, exp.Id, "")
}

func QueryExp(c *gin.Context) {
	var params request.ExperimentQueryParams
	err := c.ShouldBind(&params) // bind

	fmt.Println(params)

	if err != nil {
		response.WriteHttpErrMsg(c, 999, "exp 查询失败!", err.Error())
		return
	}
	job, err := initialize.ExpSvc.QueryExp(&params)

	if err != nil {
		response.WriteHttpErrMsg(c, 999, "exp query error!", err.Error())
		return
	}
	// logger.Infof("文件上传耗时: %s", time.Since(start))
	// utils.WriteHttpOkMsgs(c, id, "")
	response.WriteHttpOkMsgs(c, job.Id, job)
}

func QueryExps(c *gin.Context) {
	var params request.ExperimentQueryParams
	err := c.ShouldBind(&params) // bind

	if err != nil {
		response.WriteHttpErrMsg(c, 999, "exp查询失败!", err.Error())
		return
	}
	exps := initialize.ExpSvc.QueryExps(&params)

	if err != nil {
		response.WriteHttpErrMsg(c, 999, "exp保存失败!", err.Error())
		return
	}
	// logger.Infof("文件上传耗时: %s", time.Since(start))
	// utils.WriteHttpOkMsgs(c, id, "")
	response.WriteHttpOkMsgs(c, "", exps)
}

func DeleteExp(c *gin.Context) {
	var params request.ExperimentQueryParams
	err := c.ShouldBindJSON(&params) // bind

	if err != nil {
		response.WriteHttpErrMsg(c, 999, "job删除失败!", err.Error())
		return
	}
	err = initialize.ExpSvc.DeleteExp(&params)

	if err != nil {
		response.WriteHttpErrMsg(c, 999, "job删除失败!", err.Error())
		return
	}
	// logger.Infof("文件上传耗时: %s", time.Since(start))
	// utils.WriteHttpOkMsgs(c, id, "")
	response.WriteHttpOkMsgs(c, "", "")
}
