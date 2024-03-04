package main

import (
	"fmt"
	"log"
	"net"
)

func server(s string) {
	// Устанавливаем адрес сервера
	serverAddr, err := net.ResolveUDPAddr("udp", "127.0.0.1:2000")
	if err != nil {
		log.Println(err.Error())
	}
	// Создаем UDP соединение
	conn, err := net.ListenUDP("udp", serverAddr)
	if err != nil {
		log.Println(err.Error())
	}
	fmt.Println("Server started")

	broadcastAddr, err := net.ResolveUDPAddr("udp", "255.255.255.255:2002")
	if err != nil {
		log.Println(err.Error())
	}
	// Обработка входящих сообщений
	buf := make([]byte, 1024)
	fmt.Println(conn)
	for {
		n, addr, err := conn.ReadFromUDP(buf)
		if err != nil {
			log.Println(err.Error())
		}
		fmt.Printf("Received: %s from %s (%d)\n", string(buf[:n]), addr, n)

		// Отправляем сообщение всем клиентам через широковещательный адрес
		count, err := conn.WriteToUDP(buf[:n], broadcastAddr)
		if err != nil {
			log.Println(err.Error())
		}
		fmt.Println(count)
	}
}
