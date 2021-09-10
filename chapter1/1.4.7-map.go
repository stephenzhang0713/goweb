package main

import "fmt"

func main() {
	var literalMap map[string]string
	var assignedMap map[string]string
	literalMap = map[string]string{"first": "go", "second": "web"}
	createdMap := make(map[string]float32)
	assignedMap = literalMap
	createdMap["k1"] = 99
	createdMap["k2"] = 199
	assignedMap["second"] = "program"
	fmt.Printf("Map literal at \"first\" is: %s\n", literalMap["first"])
	fmt.Printf("Map created at \"k2\" is: %f\n", createdMap["k2"])
	fmt.Printf("Map assgined at \"second\" is: %s\n", literalMap["second"])
	fmt.Printf("Map literal at \"third\" is: %s\n", literalMap["third"])
}
