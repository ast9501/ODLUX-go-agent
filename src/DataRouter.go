package src

import (
	"agent/service"
	"github.com/gin-gonic/gin"
)

func AddDataRouter(r *gin.RouterGroup){
	reader := r.Group("/")

	// Mount netconf devices
	reader.PUT("network-topology:network-topology/topology=topology-netconf/node=:nodeName", service.MountDevice)
}