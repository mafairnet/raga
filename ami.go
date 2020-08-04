package main

func queueShowAll() {
	queuesData := ""
	finalCommand := login

	finalCommand = finalCommand + "Action: Queues\r\n\r\n"

	finalCommand = finalCommand + logoff

	queuesData = SocketClient(configuration.AsteriskIP, configuration.AmiPort, finalCommand)
	_ = queuesData
}

func command(command string) string {
	data := ""
	finalCommand := login

	finalCommand = finalCommand + "Action: Command\r\nCommand: " + command + "\r\n\r\n"

	finalCommand = finalCommand + logoff

	data = SocketClient(configuration.AsteriskIP, configuration.AmiPort, finalCommand)
	_ = data
	return data
}
