package model

import "encoding/json"

type OnosGetDeviceRes struct {
	// Extract raw json message at field "devices"
	Devices	json.RawMessage	`json:"devices"`
}

type OnosDevice struct {
	Name	string
	Ip		string
	Port	string
	Status	string
}

func (d *OnosDevice) GenerateReq(passwd string) (req []byte){
	payload := "{\"devices\": {\"netconf:" + d.Ip + ":" + d.Port + "\": {\"netconf\": {\"ip\":\"" + d.Ip + "\", \"port\":" + d.Port + ", \"username\": \"" + d.Name + "\", \"password\": \"" + passwd + "\" }, \"basic\": { \"driver\": \"ovs-netconf\"} } } }"
	req = []byte(payload)
	return
}