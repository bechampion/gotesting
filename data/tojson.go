package main

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	Name string `json:"name,omitempty"`
	Age  int    `json:"age,omitempty"`
}

func main() {
	jerry := Person{"Jeronimo", 13}
	var data []byte
	data, err := json.Marshal(jerry)
	_ = err
	//err := json.Unmarshal(data, &jerry)
	fmt.Println(string(data))
}
