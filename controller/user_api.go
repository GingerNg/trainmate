package controller

import (
	"trainmate/models/response"

	"github.com/gin-gonic/gin"
)

func UserLogin(c *gin.Context) {
	data := make(map[string]interface{})
	data["token"] = "admin-token"
	response.WriteHttpOkMsgs(c, "", data)
}

func UserInfo(c *gin.Context) {
	data := make(map[string]interface{})
	data["roles"] = []string{"admin"}
	data["introduction"] = "I am a super administrator"
	data["avatar"] = "https://wpimg.wallstcn.com/f778738c-e4f8-4870-b634-56703b4acafe.gif"
	data["name"] = "Super Admin"
	response.WriteHttpOkMsgs(c, "", data)
}
