package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type User struct {
	UserName string
	NickName string `json:"nick_name"`
	Email    string
}

func main() {
	user := &User{
		UserName: "Jack",
		NickName: "Ma",
		Email:    "xxxxx@qq.com",
	}

	data, err := json.Marshal(user)
	if err != nil {
		fmt.Printf("json.Marshal failed, err:%v\n", err)
		return
	}

	fmt.Printf("%s\n", string(data))

	file, _ := os.Create("chapter6/6.3/json_write.json")
	defer file.Close()
	file.Write(data)

}
