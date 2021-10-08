package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"time"
)

func main() {
	server := "127.0.0.1:8086"
	tcpAddr, err := net.ResolveTCPAddr("tcp4", server)
	if err != nil {
		Log(os.Stderr, "Fatal error:", err.Error())
		os.Exit(1)
	}
	conn, err := net.DialTCP("tcp", nil, tcpAddr)
	if err != nil {
		Log("Fatal error:", err.Error())
		return
	}
	Log(conn.RemoteAddr().String(), "connect success!")
	Sender(conn)
	Log("end")
}

func Sender(conn *net.TCPConn) {
	defer conn.Close()
	sc := bufio.NewReader(os.Stdin)
	go func() {
		t := time.NewTicker(time.Second)
		defer t.Stop()
		for {
			<-t.C
			_, err := conn.Write([]byte("1"))
			if err != nil {
				fmt.Println(err.Error())
				return
			}
		}
	}()
	name := ""
	fmt.Println("请输入聊天昵称")
	fmt.Fscan(sc, &name)
	msg := ""
	buffer := make([]byte, 1024)
	_t := time.NewTimer(time.Second * 5)
	defer _t.Stop()

	go func() {
		<-_t.C
		fmt.Println("服务器出现故障，断开连接")
		return
	}()
	for {
		go func() {
			for {
				n, err := conn.Read(buffer)
				if err != nil {
					return
				}
				_t.Reset(time.Second * 5)
				if string(buffer[0:1]) != "1" {
					fmt.Println(string(buffer[0:n]))
				}
			}
		}()
		fmt.Fscan(sc, &msg)
		i := time.Now().Format("2006-01-02 15:04:05")
		conn.Write([]byte(fmt.Sprintf("%s\n\t%s: %s", i, name, msg)))
	}
}

func Log(v ...interface{}) {
	fmt.Println(v...)
	return
}
