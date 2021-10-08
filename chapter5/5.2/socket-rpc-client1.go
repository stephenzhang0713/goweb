package main

import (
	"fmt"
	"log"
	"net/rpc/jsonrpc"
)

func main() {
	fmt.Println("client start......")
	client, err := jsonrpc.Dial("tcp", "127.0.0.1:8085")
	if err != nil {
		log.Fatal("Dial err=", err)
	}
	send := Send{"Java", "Go"}
	var receive string
	err = client.Call("Programmer.GetSkill", send, &receive)
	if err != nil {
		fmt.Println("Call err=", err)
	}
	fmt.Println("receive", receive)
}

// 参数结构体可以和服务端不一样
// 但是结构体里的字段必须一样
type Send struct {
	Java, Go string
}
