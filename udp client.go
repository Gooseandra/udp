package main

import (
	"fmt"
	"net"
)

func client(a, s, d string) {
	// Определение адреса широковещательной рассылки и порта
	serverAddr, err := net.ResolveUDPAddr("udp", ":2002")
	if err != nil {
		fmt.Println("Error resolving server address:", err)
		return
	}

	// Установка соединения с сервером
	conn, err := net.ListenUDP("udp", serverAddr)
	if err != nil {
		fmt.Println("Error connecting to server:", err)
		return
	}
	defer conn.Close()

	servAddr, err := net.ResolveUDPAddr("udp", "127.0.0.1:2000")
	if err != nil {
		fmt.Println("Error resolving server address:", err)
		return
	}

	fmt.Println("Client connected to server")
	conn.WriteToUDP([]byte("fdd"), servAddr)
	// Цикл для приема сообщений от сервера
	buffer := make([]byte, 1024)
	fmt.Println(conn)
	for {
		n, addr, err := conn.ReadFromUDP(buffer)
		if err != nil {
			fmt.Println("Error reading from server:", err)
			return
		}

		fmt.Printf("Received message from server %s: %s\n", addr, string(buffer[:n]))
	}
}
