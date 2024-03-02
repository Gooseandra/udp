package main

import (
	"fmt"
	"log"
	"net"
)

func server(port string) {
	fmt.Println("Launching server...")

	addr, err := net.ResolveUDPAddr("udp", "127.0.0.1:8080")
	if err != nil {
		fmt.Print(err)
		return
	}

	udpServer, err := net.ListenUDP("udp", addr)

	if err != nil {
		log.Fatal(err)
	}
	defer udpServer.Close()

	for {
		buf := make([]byte, 1024)
		_, addres, err := udpServer.ReadFromUDP(buf)
		if err != nil {
			fmt.Println(err.Error())
			continue
		}
		fmt.Println("Message from client: " + string(buf))
		go response(udpServer, addres, buf)
	}

}

func response(udpServer *net.UDPConn, addr *net.UDPAddr, buf []byte) {
	responseStr := fmt.Sprintf(string(buf))
	_, err := udpServer.WriteToUDP([]byte(responseStr), addr)
	if err != nil {
		fmt.Println(err.Error())
	}
	//udpServer.WriteTo([]byte(responseStr), addr)
}
