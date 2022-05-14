package pkg

import (
	. "agent/model"
	"strconv"
)

type NewNodeInfo struct {
	Host		string 		`json:"host"`
	Id			string 		`json:"id"`
	Required	bool 		`json:"is-required"`
	NodeId		string 		`json:"node-id"`
	Port		int 		`json:"port"`
	Username	string 		`json:"username"`
	Passwd		string		`json:"password"`
}

type NewDevReq struct {
	Data	NewNodeInfo	`json:"data-provider:input"`
}

func (r *NewDevReq) Parser() (device OnosDevice, passwd string) {
	device.Name = r.Data.Username
	device.Ip = r.Data.Host
	device.Port = strconv.Itoa(r.Data.Port)
	passwd = r.Data.Passwd

	return
}