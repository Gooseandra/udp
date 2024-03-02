package main

import (
	"bytes"
	"fmt"
	"log"
	"net"
)

type user struct {
	name string
	addr []byte
}

func server(port string) {
	fmt.Println("Launching server...")
	var users []user
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
	mes := make(chan []byte, 2)
	for {
		buf := make([]byte, 1024)
		count, address, err := udpServer.ReadFromUDP(buf)
		if count > 0 {
			if len(users) == 0 {
				fmt.Println("Первый клиент, добавляю в миассив")
				newUser := user{name: string(buf), addr: address.IP}
				users = append(users, newUser)
			} else {
				fmt.Println("Не первый клиент")
				for index, item := range users {
					if bytes.Equal(item.addr, address.IP) {
						fmt.Println("Знакомый")
						mes <- buf
						if err != nil {
							fmt.Println(err.Error())
							continue
						}
						fmt.Println("Message from client: " + string(buf))
						go response(udpServer, address, mes, item.name, users)
						break
					}
					if index == len(users) {
						fmt.Println("Новый клиент")
						newUser := user{name: string(buf), addr: address.IP}
						users = append(users, newUser)
					}
				}
			}
		}
	}
}

func response(udpServer *net.UDPConn, addr *net.UDPAddr, mes chan []byte, name string, users []user) {
	responseStr := fmt.Sprintf(string(<-mes))
	for _, item := range users {
		newAddress := net.UDPAddr{IP: item.addr, Port: 8080}
		_, err := udpServer.WriteToUDP([]byte(responseStr), &newAddress)
		if err != nil {
			fmt.Println(err.Error())
		}
	}
	//udpServer.WriteTo([]byte(responseStr), addr)
}
