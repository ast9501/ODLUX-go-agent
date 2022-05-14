package internal

import (
	"os"
	. "agent/model"
	"net/http"
	//"encoding/json"
	"log"
	"io/ioutil"
	"time"
	. "agent/pkg"
	"bytes"
)

/*
type onosServer struct {
	*OnosServer
}
*/
type OnosServer struct {
	Ip			string
	Port		string
	Username	string
	Password	string
}

func (s *OnosServer) CreateDevice(device OnosDevice, passwd string) (status bool){
	
	url := "http://" + s.Ip + ":" + s.Port + "/onos/v1/network/configuration"

	client := &http.Client{
		Timeout: time.Second * 10,
	}
	

	body := device.GenerateReq(passwd)
	reqBody := bytes.NewReader(body)
	req, err := http.NewRequest("POST", url, reqBody)
	if err != nil {
		log.Println("CreateDevice() generate Http request error")
		return false
	}

	req.SetBasicAuth(s.Username, s.Password)
	req.Header.Set("Content-Type", "application/json")
	res, err := client.Do(req)
	
	if err != nil {
		log.Println("CreateDevice(): call ONOS api error! ")
		return false
	}
	defer res.Body.Close()

	if (res.StatusCode != 200){
		log.Println("CreateDevice(): Unexpected status code ", res.StatusCode, " from onos response")
		return false
	}
	return true
}

func (s *OnosServer) GetDevices() (deviceList interface{}) {

	url := "http://" + s.Ip + ":" + s.Port +"/onos/v1/network/configuration"

	// TODO: Decouple function: Raise request
	client := &http.Client{
		Timeout: time.Second * 10,
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Println("GetDevices() generate Http Request error")
	}
	req.SetBasicAuth(s.Username, s.Password)
	res, err := client.Do(req)

	if err != nil {
		log.Println(err)
	}
	defer res.Body.Close()

	data, err := ioutil.ReadAll(res.Body)

	if err != nil {
		log.Println(err)
		return
	}
	deviceList = ParsOnosResponse(data, "GetDevice")

	return 
}

func InitOnos() *OnosServer {
	onos := getOnosVar()
	return onos
}

func getOnosVar() *OnosServer {
	onos := &OnosServer{}
	onos.Ip = os.Getenv("ONOS_IP")
	if (isExist(os.Getenv("ONOS_PORT"))) {
		onos.Port = os.Getenv("ONOS_PORT")
	} else {
		onos.Port = "8181"
	}

	if (isExist(os.Getenv("ONOS_USER"))) {
		onos.Username = os.Getenv("ONOS_USER")
	} else {
		onos.Username = "onos"
	}
	
	if (isExist(os.Getenv("ONOS_PASSWORD"))) {
		onos.Password = os.Getenv("ONOS_PASSWORD")
	} else {
		onos.Password = "rocks"
	}
	
	return onos
}

func isExist(env string) bool {
	if (env != "") {
		return true
	}
	return false
}