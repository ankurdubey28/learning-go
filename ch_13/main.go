package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	p := Person{
		Name: "ankur",
		Age:  22,
	}
	out, err := json.Marshal(p)
	if err == nil {
		fmt.Println(out)
	}
}

type Person struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}
