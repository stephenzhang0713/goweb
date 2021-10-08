package main

import (
	"fmt"
	"net"
	"time"
)

var ConnSlice map[net.Conn]*Heartbeat

type Heartbeat struct {
	endTime int64 //过期时间
}

func main() {
	ConnSlice = map[net.Conn]*Heartbeat{}
	l, err := net.Listen("tcp", "127.0.0.1:8086")
	if err != nil {
		fmt.Println("服务器启动失败")
	}
	defer l.Close()
	for {
		conn, err := l.Accept()
		if err != nil {
			fmt.Println("Error Accepting:", err)
		}
		fmt.Printf("Received message %s -> %s \n", conn.RemoteAddr(), conn.LocalAddr())
		ConnSlice[conn] = &Heartbeat{endTime: time.Now().Add(time.Second * 5).Unix()}
		go handleConn(conn)
	}
}

func handleConn(c net.Conn) {
	buffer := make([]byte, 1024)
	for {
		n, err := c.Read(buffer)
		if ConnSlice[c].endTime > time.Now().Unix() {
			ConnSlice[c].endTime = time.Now().Add(time.Second * 5).Unix()
		} else {
			fmt.Println("长时间未发消息断开连接")
			return
		}
		if err != nil {
			return
		}
		if string(buffer[0:n]) == "1" {
			c.Write([]byte("1"))
			continue
		}
		for conn, heart := range ConnSlice {
			if conn == c {
				continue
			}
			if heart.endTime < time.Now().Unix() {
				delete(ConnSlice, conn)
				conn.Close()
				fmt.Println("删除连接", conn.RemoteAddr())
				fmt.Println("现在存有连接", ConnSlice)
				continue
			}
			conn.Write(buffer[0:n])
		}
	}
}
