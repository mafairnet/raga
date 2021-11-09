package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
)

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "RAGA | REST Asterisk Golang Agent, please look the API documentation for the supported requests!\n")
}

func Command(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	vars := mux.Vars(r)

	commandString := vars["command"]
	token := r.URL.Query().Get("token")

	if token == configuration.Token {
		data := command(commandString)

		if err := json.NewEncoder(w).Encode(data); err != nil {
			panic(err)
		}
	} else {
		fmt.Fprint(w, "{\"error\":\"Not Authorized\"}")
	}
}

func Monitor(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	token := r.URL.Query().Get("token")

	commandString := "queue show"

	if token == configuration.Token {
		//Get raw data
		data := command(commandString)

		dataArray := strings.Split(data, "\n")

		//fmt.Printf("Data: %v \n", dataArray)

		var queuesArray []string
		var agentsArray []Agent
		var callerArray []string
		var queues []QueueData
		var finalQueue QueueData
		queueFound := false
		queueIndex := 0
		currentQueue := ""

		//Get queues
		for _, lineRaw := range dataArray {

			if strings.Contains(lineRaw, "strategy") {
				if !strings.Contains(lineRaw, "default") {
					queueRaw := strings.Split(lineRaw, " ")
					currentQueue = queueRaw[0]
					//fmt.Printf("QueueRaw: %v\n", queue)
					queuesArray = append(queuesArray, currentQueue)
					queueFound = true
					queueIndex = queueIndex + 1
				}
			}

			if queueFound {
				if strings.Contains(lineRaw, "has taken") {
					if strings.Contains(lineRaw, "from-internal") {
						memberRaw := strings.Split(lineRaw, "@")
						memberRaw = strings.Split(memberRaw[0], "/")
						member := memberRaw[1]
						member = strings.Replace(member, " ", "", -1)
						status := "unavailable"
						if strings.Contains(lineRaw, "Not in use") {
							status = "available"
						}
						if strings.Contains(lineRaw, "In use") {
							status = "busy"
						}
						if strings.Contains(lineRaw, "On Hold") {
							status = "busy"
						}
						if strings.Contains(lineRaw, "paused") {
							status = "paused"
						}
						if strings.Contains(lineRaw, "Unavailable") {
							status = "unavailable"
						}
						agent := Agent{member, status}
						agentsArray = append(agentsArray, agent)
					}
				}

				if strings.Contains(lineRaw, "wait: ") && strings.Contains(lineRaw, "prio: ") {
					callerRaw := strings.Split(lineRaw, " (wait")
					caller := callerRaw[0]
					caller = strings.Replace(caller, " ", "", -1)
					//caller := strings.Replace(callerRaw[0], "SIP", "", -1)
					callerArray = append(callerArray, caller)
				}

				if lineRaw == "" {
					finalQueue = QueueData{currentQueue, agentsArray, callerArray}
					queues = append(queues, finalQueue)
					callerArray = nil
					agentsArray = nil
					//queueFound = false
				}
			}
		}

		fmt.Printf("Queues: %v\n", queuesArray)
		//fmt.Printf("Agents: %v\n", agentsArray)

		fmt.Printf("FinalQueues: %v\n", queues)

		//Get queue calls

		if err := json.NewEncoder(w).Encode(queues); err != nil {
			panic(err)
		}
	} else {
		fmt.Fprint(w, "{\"error\":\"Not Authorized\"}")
	}
}

func Devices(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	commandString := "sip show peers"

	var devices []Device
	_ = devices

	if configuration.PjSipEnabled {
		commandString = "pjsip show endpoints"
	}

	token := r.URL.Query().Get("token")

	if token == configuration.Token {
		data := command(commandString)

		devices = processDevices(data)

		if err := json.NewEncoder(w).Encode(devices); err != nil {
			panic(err)
		}
	} else {
		fmt.Fprint(w, "{\"error\":\"Not Authorized\"}")
	}
}

func QueueAction(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	vars := mux.Vars(r)

	action := vars["action"]
	extension := r.URL.Query().Get("extension")
	queue := r.URL.Query().Get("queue")
	token := r.URL.Query().Get("token")

	if token == configuration.Token {
		data := ""

		switch action {
		case "login":
			data = AmiLoginQueue(extension, queue)
			_ = data
		case "logout":
			data = AmiLogoutQueue(extension, queue)
			_ = data
		case "pause":
			data = AmiPauseQueue(extension, queue)
			_ = data
		case "unpause":
			data = AmiUnPauseQueue(extension, queue)
			_ = data
		}

		message := "{\"message\":\"" + data + "\"}"

		if err := json.NewEncoder(w).Encode(message); err != nil {
			panic(err)
		}
	} else {
		fmt.Fprint(w, "{\"error\":\"Not Authorized\"}")
	}
}

func wsEndpoint(w http.ResponseWriter, r *http.Request) {
	upgrader.CheckOrigin = func(r *http.Request) bool { return true }

	// upgrade this connection to a WebSocket
	// connection
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
	}
	// helpful log statement to show connections
	log.Println("Client Connected")

	err = ws.WriteMessage(1, []byte("Hi Client!"))
	if err != nil {
		log.Println(err)
	}

	//reader(ws)
	commandReader(ws)
}
