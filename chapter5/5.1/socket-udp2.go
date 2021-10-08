package main

import (
	"fmt"
	"net"
)

func main() {
	laddr, err := net.ResolveUDPAddr("udp", "127.0.0.1:8023")
	if err != nil {
		fmt.Println("ResolveUDPAddr err:", err)
		return
	}
	conn, err := net.ListenUDP("udp", laddr)
	if err != nil {
		fmt.Println("net.ListenUDP err:", err)
		return
	}
	defer conn.Close()

	for {
		buf := make([]byte, 1024)
		n, raddr, err := conn.ReadFromUDP(buf)
		if err != nil {
			fmt.Println("conn.ReadFromUDP err:", err)
			return
		}
		fmt.Printf("接受到客户端[%s]:%s", raddr, string(buf[:n]))

		conn.WriteToUDP([]byte("i am server"), raddr)
	}
}
