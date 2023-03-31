package utils

import (
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func FileExists(path string) bool {
	_, err := os.Stat(path)
	if err == nil {
		return true
	}
	if os.IsNotExist(err) {
		return false
	}
	return false
}

type CommonRes struct {
	Version      string `json:"version"`
	Status       string `json:"status"`
	Method       string `json:"method"`
	ResCode      string `json:"code"`
	ResMsg       string `json:"msg"`
	Id           string `json:"id"`
	Data         interface{}
	ErrorDetails interface{} `json:"error_details,omitempty"`
}

func WriteHttpOkMsgs(c *gin.Context, id string, v interface{}) {
	ok := CommonRes{
		Status:  "1",
		Version: "v1.0",
		Method:  c.Request.URL.Path,
		ResCode: "0000",
		ResMsg:  "success",
		Id:      id,
		Data:    v,
	}

	c.JSON(http.StatusOK, ok)
}

func WriteFileOk(c *gin.Context, contentType string, data []byte) {
	c.Data(http.StatusOK, contentType, data)
}

func WriteHttpErrMsg(c *gin.Context, code string, msg string, details ...interface{}) {
	err := CommonRes{
		Status:       "0",
		Version:      "v1.0",
		Method:       c.Request.URL.Path,
		ResCode:      code,
		ResMsg:       msg,
		ErrorDetails: details,
	}
	c.Header("Parsed-Status", "500-100")
	// logger.Infof("[req-%s]%+v", c.Request.Header.Get(constants.HttpRequestIdKey), err)
	c.JSON(http.StatusOK, err)
}
