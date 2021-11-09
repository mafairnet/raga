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

func AmiLoginQueue(extension string, queue string) string {
	data := ""
	finalCommand := login

	command := "queue add member Local/" + extension + "@from-internal/n to " + queue + " penalty 0 as \"Local/" + extension + "@from-internal/n\" state_interface SIP/" + extension

	finalCommand = finalCommand + "Action: Command\r\nCommand: " + command + "\r\n\r\n"

	finalCommand = finalCommand + logoff

	data = SocketClient(configuration.AsteriskIP, configuration.AmiPort, finalCommand)
	_ = data
	return data
}

func AmiLogoutQueue(extension string, queue string) string {
	data := ""
	finalCommand := login

	command := "queue remove member Local/" + extension + "@from-internal/n from " + queue

	finalCommand = finalCommand + "Action: Command\r\nCommand: " + command + "\r\n\r\n"

	finalCommand = finalCommand + logoff

	data = SocketClient(configuration.AsteriskIP, configuration.AmiPort, finalCommand)
	_ = data
	return data
}

func AmiPauseQueue(extension string, queue string) string {
	data := ""
	finalCommand := login

	command := "queue pause member Local/" + extension + "@from-internal/n queue " + queue

	finalCommand = finalCommand + "Action: Command\r\nCommand: " + command + "\r\n\r\n"

	finalCommand = finalCommand + logoff

	data = SocketClient(configuration.AsteriskIP, configuration.AmiPort, finalCommand)
	_ = data
	return data
}

func AmiUnPauseQueue(extension string, queue string) string {
	data := ""
	finalCommand := login

	command := "queue unpause member Local/" + extension + "@from-internal/n queue " + queue

	finalCommand = finalCommand + "Action: Command\r\nCommand: " + command + "\r\n\r\n"

	finalCommand = finalCommand + logoff

	data = SocketClient(configuration.AsteriskIP, configuration.AmiPort, finalCommand)
	_ = data
	return data
}
