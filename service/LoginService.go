package service

import (
	//"os"
	"net/http"
	"github.com/gin-gonic/gin"
	. "agent/internal"
	//. "agent/model"
	//. "agent/internal"
	"encoding/json"
	"log"
)

func GetProvider (c *gin.Context) {
	data := LoadOauthProvider()
	if len(data) == 0 {
		c.JSON(http.StatusInternalServerError, "Missing config")
		return
	}

	c.Data(http.StatusOK, "application/json", data)
	return
}

func VerifyUser (c *gin.Context) {
	user := c.PostForm("username")
	passwd := c.PostForm("password")

	if (user == "admin" && passwd == "admin"){
		token := GenerateToken()
		tokenJson, _ := json.Marshal(token)
		log.Println(string(tokenJson))
		c.Data(http.StatusOK, "application/json", tokenJson)
		//c.JSON(http.StatusOK, tokenJson)
		//c.JSON(http.StatusOK, token)
		return
	}
}