package main

import (
	"bufio"
	"net"
	"os"
)

func writer(conn *net.UDPConn, mes chan string) {
	for {
		reader := bufio.NewReader(os.Stdin)
		text, _ := reader.ReadString('\n')
		if text == "exit" {
			mes <- text
			return
		}
		conn.Write([]byte(text))
	}
}

func listener(conn *net.UDPConn) {
	for {
		//message, err := bufio.NewReader(conn).ReadString('\n')
		//if err != nil {
		//	fmt.Println(err.Error())
		//	return
		//}
		//fmt.Println(message)
		received := make([]byte, 1024)
		_, err := conn.Read(received)
		if err != nil {
			println("Read data failed:", err.Error())
			os.Exit(1)
		}
	}
}

func client(ip, port, name string) {
	udpServer, err := net.ResolveUDPAddr("udp", ":8080")

	if err != nil {
		println("ResolveUDPAddr failed:", err.Error())
		os.Exit(1)
	}

	conn, err := net.DialUDP("udp", nil, udpServer)
	if err != nil {
		println("Listen failed:", err.Error())
		os.Exit(1)
	}

	var exit chan string
	go listener(conn)
	go writer(conn, exit)
	defer conn.Close()
	<-exit
	// buffer to get data

	//received := make([]byte, 1024)
	//_, err = conn.Read(received)
	//if err != nil {
	//	println("Read data failed:", err.Error())
	//	os.Exit(1)
	//}
	//
	//println(string(received))
}
