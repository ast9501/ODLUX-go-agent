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