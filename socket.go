package main

import (
	"io/ioutil"
	"log"
	"net"
	"strconv"
	"strings"
)

func SocketClient(ip string, port int, message string) string {

	socketResponse := ""
	const StopCharacter = "\r\n\r\n"
	messageInBytes := []byte(message)
	addr := strings.Join([]string{ip, strconv.Itoa(port)}, ":")
	conn, err := net.Dial("tcp", addr)

	if err != nil {
		log.Fatalln(err)
	} else {

		defer conn.Close()

		conn.Write([]byte(messageInBytes))
		conn.Write([]byte(StopCharacter))

		log.Printf("Send: %s", message)

		var response, _ = ioutil.ReadAll(conn)
		log.Printf("Receive: %s", response)
		socketResponse = string(response)
	}

	return socketResponse
}
