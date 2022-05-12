package pkg

import (
	. "agent/model"
	"encoding/json"
	"log"
	//"strconv"
	"fmt"
)

func ResponseBuilder(data interface{}) (response []byte) {
	switch devices := data.(type) {
	case []OnosDevice:
		provider := DataProviderOutput{}
		for _, device := range(devices) {
			// TODO: Get device Capacity from ONOS
			caps := NodeDtail{
				UnavailCap: []string{"(urn:sysrepo:plugind?revision=2020-12-10)sysrepo-plugind"},
				AvailCap:	[]string{"(urn:o-ran:dhcp:1.0?revision=2020-12-10)o-ran-dhcp"},
			}
			tmp := NodeInfo{
				CoreMdlCap: "Unsupported",
				Type:		"O-RAN",
				Id:			device.Name,
				Dtail:		caps,
				Required:	false,
				Host:		device.Ip,
				Port:		device.Port,
				NodeId:		device.Name,
				Status:		device.Status,
			}
			provider.Data = append(provider.Data, tmp)
		}
		provider.Page = Pagination {
			Size:	10,
			Page:	"1",
			Total:	len(provider.Data),
		}
		resBody := OdlReadConnRes{
			Body:	provider,
		}
		response, _ = json.Marshal(resBody)
		return
	// another case
	}
	return
}

func ParsOnosResponse(data []byte, response string) (result interface{}) {
	//var result interface{}
	switch response {
	case "GetDevice":
		var resBody OnosGetDeviceRes
		// Parsing json with non-static field
		var deviceList interface{}
		json.Unmarshal(data, &resBody)
		var deviceInfoList []OnosDevice

		// Extract devices field to parse indivisual devices info
		json.Unmarshal(resBody.Devices, &deviceList)
		m := deviceList.(map[string]interface{})
		for k, v := range m {
			switch vv := v.(type) {
			case string:
				log.Println(k, " is string ", vv)
			case map[string]interface{}:
				device := OnosDevice{}
				device.Name = k
				netconfs := vv["netconf"].(map[string]interface{})
				switch vvv := netconfs["ip"].(type) {
				case string:
					device.Ip = vvv
				case float64:
					device.Ip = fmt.Sprintf("%g", vvv)
				}
				switch vvv := netconfs["port"].(type) {
				case string:
					device.Port = vvv
				case float64:
					device.Port = fmt.Sprintf("%g", vvv)
				}
				deviceInfoList = append(deviceInfoList, device)
			default:
				log.Println("Unknown type")
			}
		}
		result = deviceInfoList
		return
	}
	return
}