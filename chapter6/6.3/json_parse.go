package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

type MongoConfig struct {
	MongoAddr       string
	MongoPoolLimit  int
	MongoDb         string
	MongoCollection string
}

type Config struct {
	Port  string
	Mongo MongoConfig
}

type JsonStruct struct {
}

func NewJsonStruct() *JsonStruct {
	return &JsonStruct{}
}

func (js *JsonStruct) Load(fileName string, v interface{}) {
	data, err := ioutil.ReadFile(fileName)
	if err != nil {
		return
	}

	err = json.Unmarshal(data, v)
	if err != nil {
		return
	}
}

func main() {
	jsonParse := NewJsonStruct()
	v := Config{}
	jsonParse.Load("chapter6/6.3/json_parse.json", &v)
	fmt.Println(v.Port)
	fmt.Println(v.Mongo.MongoDb)
}
