package model

type DeviceLog struct {
	DeviceId	string `json: "deviceId"`
	EventType	string `json: "eventType"`
	TimeStamp	string `json: "timestamp"`
}