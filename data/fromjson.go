package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

type Person struct {
	Name     string
	Age      int
	LastName string
	Foods    []string
}

func main() {
	var dada []string
	file, e := ioutil.ReadFile("./config.json")
	if e != nil {
		fmt.Printf("File error: %v\n", e)
		os.Exit(1)
	}
	dada = append(dada, "1")
	fmt.Println(dada)
	//fmt.Printf("%s\n", string(file))

	var me Person
	json.Unmarshal(file, &me)
	fmt.Printf("Results: %v\n", me)
}
