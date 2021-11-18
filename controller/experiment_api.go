package controller

import (
	"trainmate/models/request"
	"trainmate/models/response"
	Service "trainmate/service"

	"github.com/gin-gonic/gin"
)

func CreateExp(c *gin.Context) {

	var params request.ExperimentParams
	err := c.ShouldBindJSON(&params) // bind

	if err != nil {
		response.WriteHttpErrMsg(c, 999, "exp保存失败!", err.Error())
		return
	}
	exp, err := Service.CreateExp(&params)

	if err != nil {
		response.WriteHttpErrMsg(c, 999, "exp保存失败!", err.Error())
		return
	}
	// logger.Infof("文件上传耗时: %s", time.Since(start))
	// utils.WriteHttpOkMsgs(c, id, "")
	response.WriteHttpOkMsgs(c, exp.Id, "")
}

func QueryExps(c *gin.Context) {
	var params request.ExperimentQueryParams
	err := c.ShouldBind(&params) // bind

	if err != nil {
		response.WriteHttpErrMsg(c, 999, "exp查询失败!", err.Error())
		return
	}
	exps := Service.QueryExp(&params)

	if err != nil {
		response.WriteHttpErrMsg(c, 999, "exp保存失败!", err.Error())
		return
	}
	// logger.Infof("文件上传耗时: %s", time.Since(start))
	// utils.WriteHttpOkMsgs(c, id, "")
	response.WriteHttpOkMsgs(c, "", exps)
}
