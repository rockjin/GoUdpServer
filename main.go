package main

import (
	"fmt"
	"log"

	"fsu.udp"
)

func main() {
	fmt.Println("Server started..")
	log.Println("Server started...")
	UdpServer.Init()
}
