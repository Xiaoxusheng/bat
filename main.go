package main

import "bat/message"

func main() {
	bat := message.NewChatBat(0.1, 1024, 1024)
	bat.Start()
}
