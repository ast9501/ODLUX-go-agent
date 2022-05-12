package service

import (
	"log"
	"net/http"
	//"strconv"
	. "agent/internal"
	. "agent/pkg"

	"github.com/gin-gonic/gin"
)


func Handler(c *gin.Context) {
	op := c.Param("operation")
	log.Println(op)
	switch op {
		case ":read-network-element-connection-list":
			GetConnectDevice(c)
		case ":create-network-element-connection":
			CreateDevice(c)
		case ":read-connectionlog-list":
			ReadDeviceLog(c)
		default:
			log.Println("Invalid Request")
	}
}

func MountDevice(c *gin.Context) {
	deviceName := c.Param("nodeName") //strconv.Atoi(c.Param("nodeName"))
	log.Println(deviceName)
	c.JSON(http.StatusOK, "OK")
	return
}

func GetConnectDevice(c *gin.Context) {
	var m map[string]interface{}
	err := c.Bind(&m)
    if err != nil {
        return
    }

	onos := InitOnos()
	deviceList := onos.GetDevices()
	response := ResponseBuilder(deviceList)

	c.Data(http.StatusOK, "application/json", response)
}

func CreateDevice(c *gin.Context) {
	var m map[string]interface{}
    err := c.Bind(&m)
    if err != nil {
        return
    }

	log.Printf("%v \n", m)
	c.JSON(http.StatusOK, "OK")
}