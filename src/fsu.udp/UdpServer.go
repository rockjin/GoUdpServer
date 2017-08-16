package UdpServer

import (
	"fmt"
	"log"
	"net"
	"os"
)

func Init() {
	fmt.Println("init udp server")
	service := ":5000"
	udpAddr, err := net.ResolveUDPAddr("udp4", service)
	checkErr(err)
	conn, err := net.ListenUDP("udp4", udpAddr)
	checkErr(err)
	handleClient(conn)
}

func checkErr(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error:%s", err.Error())
		os.Exit(-1)
	}
}

func handleClient(conn *net.UDPConn) {
	defer conn.Close()
	buff := make([]byte, 1024)
	for {
		fmt.Println("read waiting")
		len, peer, err := conn.ReadFromUDP(buff)
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		fmt.Println(peer)
		log.Println(string(buff[0:len]))
		conn.WriteToUDP([]byte("ok"), peer)
	}
}
