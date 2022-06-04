package service

import (
	//"agent/model"
	"log"
	"net/http"
	//"strconv"

	"github.com/gin-gonic/gin"
)

func ReadDeviceLog(c *gin.Context) {
	var m map[string]interface{}
    err := c.Bind(&m)
    if err != nil {
        return
    }

	log.Printf("%v \n", m)
	c.JSON(http.StatusOK, "OK")
}

func GetDeviceLog(c *gin.Context) {
	var m map[string]interface{}
    err := c.Bind(&m)
    if err != nil {
        return
    }

	log.Printf("%v \n", m)
	c.JSON(http.StatusOK, "OK")
}

func PostDeviceLog(c *gin.Context) {
	var m map[string]interface{}
    err := c.Bind(&m)
    if err != nil {
        return
    }

	log.Printf("%v \n", m)
	c.JSON(http.StatusOK, "OK")
}

func GetStatus(c *gin.Context) {
	c.JSON(http.StatusOK, "Service is OK")
}