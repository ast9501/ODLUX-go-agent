package internal

import (
	"os"
	. "agent/model"
	"net/http"
	"encoding/json"
	"log"
	"io/ioutil"
	"time"
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


func (s *OnosServer) GetDevices() []Device {
	// Declare slice []Device
	var devices []Device

	url := "http://" + s.Ip + ":" + s.Port +"/onos/v1/network/configuration"
	log.Println("Request url: ", url)

	// TODO: Decouple function: Raise request
	client := &http.Client{
		Timeout: time.Second * 10,
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Println("Generate Http Request error()")
	}
	req.SetBasicAuth(s.Username, s.Password)
	res, err := client.Do(req)

	if err != nil {
		log.Println(err)
	}
	defer res.Body.Close()

	data, err := ioutil.ReadAll(res.Body)

	// TODO: Parse JSON body
	if err != nil {
		log.Println(err)
		return devices
	}

	var resBody OnosGetDeviceRes
	// Parsing json with non-static field
	var deviceList interface{}
	json.Unmarshal(data, &resBody)

	// Extract devices field to parse indivisual devices info
	json.Unmarshal(resBody.Devices, &deviceList)
	//log.Println(deviceList)
	m := deviceList.(map[string]interface{})
	for k, v := range m {
		switch vv := v.(type) {
		case string:
			log.Println(k, " is string ", vv)
		case map[string]interface{}:
			//log.Println(k, " is map[string]interface{} ", vv)
			log.Println("Device: ", k)
			netconfs := vv["netconf"].(map[string]interface{})
			log.Println("ip: ", netconfs["ip"])
			log.Println("port: ", netconfs["port"])
		default:
			log.Println("Unknown type")
		}
	}

	// TODO: Append into devices then return 
	return devices
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