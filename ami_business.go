package main

import (
	"encoding/json"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

func processRequest(request string) string {
	var wsRequest WsRequest
	result := "{\"status\":\"ack\"}"

	json.Unmarshal([]byte(request), &wsRequest)
	fmt.Printf("RequestType: %s, RequestCommand: %s", wsRequest.Type, wsRequest.Value)

	if wsRequest.Type == "command" {
		switch wsRequest.Value {
		case "devices":
			commandString := "sip show peers"
			var devices []Device
			_ = devices

			if configuration.PjSipEnabled {
				commandString = "pjsip list endpoints"
			}
			data := command(commandString)
			devices = processDevices(data)

			devicesJson, _ := json.Marshal(devices)
			result = string(devicesJson)

		case "monitor":
			result = "{\"status\":\"not_available\"}"
		}
	}

	return result
}

func processDevices(data string) []Device {
	var devices []Device
	_ = devices

	if !configuration.PjSipEnabled {

		dataArray := strings.Split(data, "\n")
		//fmt.Printf("Data: %v \n", dataArray)

		for _, lineRaw := range dataArray {
			if strings.Contains(lineRaw, "OK") || strings.Contains(lineRaw, "UNKNOWN") || strings.Contains(lineRaw, "LAGGED") {
				space := regexp.MustCompile(`\s+`)
				lineRaw := space.ReplaceAllString(lineRaw, " ")
				//fmt.Printf("Data: %v \n", lineRaw)
				extensionArray := strings.Split(lineRaw, " ")
				extension := ""
				if strings.Contains(extensionArray[0], "/") {
					extensionData := strings.Split(extensionArray[0], "/")
					extension = extensionData[0]
				} else {
					extension = extensionArray[0]
				}

				ip := extensionArray[1]
				ip = strings.Replace(ip, "(", "", -1)
				ip = strings.Replace(ip, ")", "", -1)
				ip = strings.ToLower(ip)
				port, err := strconv.Atoi(extensionArray[6])
				_ = err
				status := strings.ToLower(extensionArray[7])
				newDevice := Device{extension, ip, port, status}
				//fmt.Printf("Data: %v \n", newDevice)
				devices = append(devices, newDevice)
			}
		}
	} else {

		data = strings.Replace(data, "Output:", "", -1)

		data = strings.Replace(data, "Endpoint:", "&Endpoint:", -1)

		//fmt.Printf("Data %v", data)

		dataArray := strings.Split(data, "&")
		//fmt.Printf("Data %v", dataArray[1])

		for _, endpointItem := range dataArray {

			if (strings.Contains(endpointItem, "Not in use") || strings.Contains(endpointItem, "In use") || strings.Contains(endpointItem, "On Hold") || strings.Contains(endpointItem, "Unavailable") || strings.Contains(endpointItem, "Avail") || strings.Contains(endpointItem, "Available")) && !strings.Contains(endpointItem, "anonymous") {

				//fmt.Printf("\nENDPOINT DATA\n")

				//fmt.Printf("ITEM: \n[%v]\n", endpointItem)

				endpointData := strings.Split(endpointItem, "\r")

				//fmt.Printf("RawData: \n[%v]\n", endpointData)

				//endpointData := strings.Split(endpointItem, "\n")

				endpoint := ""
				_ = endpoint
				ip := ""
				_ = ip
				port := 0
				_ = port
				status := ""
				_ = status
				var err error
				_ = err

				for _, endpointLine := range endpointData {

					//fmt.Printf("Item: %v\n", endpointLine)

					if strings.Contains(endpointLine, "Endpoint:") {
						endpointLineRaw := strings.Split(endpointLine, "/")[1]
						space := regexp.MustCompile(`\s+`)
						endpointLineData := space.ReplaceAllString(endpointLineRaw, " ")
						//fmt.Printf("Data: %v \n", lineRaw)
						endpointArray := strings.Split(endpointLineData, " ")
						endpoint = endpointArray[0]
						status = endpointArray[1]
						//fmt.Printf("Endpoint: %v\n", endpoint)
					}
					if strings.Contains(endpointLine, "Contact:") {
						endpointLineRaw := strings.Split(endpointLine, "@")[1]
						endpointLineRaw = strings.Split(endpointLineRaw, ";")[0]
						endpointContactInfo := strings.Split(endpointLineRaw, ":")
						ip = endpointContactInfo[0]
						port, err = strconv.Atoi(endpointContactInfo[1])
						//fmt.Printf("Endpoint: %v\n", status)
					}
				}

				newDevice := Device{endpoint, ip, port, status}

				//fmt.Printf("Device: %v\n", newDevice)

				devices = append(devices, newDevice)
			}
		}
	}
	return devices
}
