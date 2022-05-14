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
	//var m map[string]interface{}
	m := NewDevReq{}
    err := c.BindJSON(&m)
	if err != nil {
		log.Println(err)
        return
    }
	deviceInfo, devicePasswd := m.Parser()

	onos := InitOnos()
	status := onos.CreateDevice(deviceInfo, devicePasswd)
	if (status) {
		c.JSON(http.StatusOK, "OK")	
		//log.Println("(ONOS) create device ", deviceInfo, " success!")
	} else {
		c.JSON(http.StatusInternalServerError, "Error: Please check your ONOS connection or verfiy your device info is correct")
	}
}