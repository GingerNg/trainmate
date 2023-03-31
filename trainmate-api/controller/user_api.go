package controller

import (
	"net/http"
	"trainmate/models/response"
	"trainmate/service"

	"github.com/gin-gonic/gin"
)

func SayHello(c *gin.Context) {
	strToken := c.Param("token")
	claim, err := service.VerifyAction(strToken)
	if err != nil {
		c.String(http.StatusNotFound, err.Error())
		return
	}
	c.String(http.StatusOK, "hello,", claim.Username)
}

func UserLogin(c *gin.Context) {
	respData := make(map[string]interface{})
	reqParams := make(map[string]interface{})
	c.BindJSON(&reqParams)
	signedToken, err := service.Login(reqParams["username"].(string), reqParams["password"].(string))
	if err != nil {
		// c.String(http.StatusNotFound, err.Error())
		// data
		response.WriteHttpErrMsg(c, 999, "err", respData)
	}
	// data["token"] = "admin-token"
	respData["token"] = signedToken
	response.WriteHttpOkMsgs(c, "", respData)
}

func RefreshToken(c *gin.Context) {

}

func UserLogout(c *gin.Context) {
	respData := make(map[string]interface{})
	respData["roles"] = []string{"admin"}
	respData["introduction"] = "I am a super administrator"
	respData["avatar"] = "https://wpimg.wallstcn.com/f778738c-e4f8-4870-b634-56703b4acafe.gif"
	response.WriteHttpOkMsgs(c, "", respData)
}

func UserInfo(c *gin.Context) {
	respData := make(map[string]interface{})
	strToken := c.Query("token") // 获取url参数
	// fmt.Println(strToken)
	claims, err := service.VerifyAction(strToken)
	if err != nil {
		// c.String(http.StatusNotFound, err.Error())
		// data
		response.WriteHttpErrMsg(c, 999, "err", respData)
	}
	respData["roles"] = []string{"admin"}
	respData["introduction"] = "I am a super administrator"
	respData["avatar"] = "https://wpimg.wallstcn.com/f778738c-e4f8-4870-b634-56703b4acafe.gif"
	// respData["name"] = "Super Admin"
	respData["name"] = claims.FullName
	response.WriteHttpOkMsgs(c, "", respData)
}
