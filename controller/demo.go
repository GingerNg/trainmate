package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// HelloWorld godoc
// @Summary ping example
// @Schemes
// @Description do ping
// @Tags example
// @Accept json
// @Produce json
// @Success 200 {string} Helloworld
// @Router /demo/hello [get]
func Helloworld(c *gin.Context) {
	c.JSON(http.StatusOK, "helloworld")
}

func UserName(c *gin.Context) {
	name := c.Param("name")
	c.String(http.StatusOK, "Hello %s", name)
}

// UserNameAction godoc
// @Summary UserNameAction
// @Schemes
// @Description do ping
// @Tags example
// @Accept json
// @Produce json
// @Success 200 {string} Helloworld
// @Router /demo/user/name/action [get]
func UserNameAction(c *gin.Context) {
	name := c.Param("name")
	action := c.Param("action")
	message := name + " is " + action
	c.String(http.StatusOK, message)
}

func Welcome(c *gin.Context) {
	firstname := c.DefaultQuery("firstname", "Guest")
	lastname := c.Query("lastname")
	c.String(http.StatusOK, "Hello %s %s", firstname, lastname)

}

// FormPost godoc
// @Summary FormPost
// @Schemes
// @Description FormPost demo
// @Tags example
// @Accept json
// @Produce json
// @Success 200 {string} Helloworld
// @Router /demo/formpost [post]
func FormPost(c *gin.Context) {
	message := c.PostForm("message")
	nick := c.DefaultPostForm("nick", "anonymous")

	c.JSON(http.StatusOK, gin.H{
		"status": gin.H{
			"status_code": http.StatusOK,
			"status":      "ok",
		},
		"message": message,
		"nick":    nick,
	})
}
