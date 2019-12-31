package main

import (
	"encoding/json"
	"fmt"
	"k8s-operator-learning/go-learning/slice"
)

type Bar struct {
	Title string
	Name  string
}
type Foo struct {
	Bar    Bar `json:",inline"`
	Number int `json:"number1"`
}

func main() {
	fmt.Println("Hello world")
	slice.SliceDemo()
	slice.SliceDemo1()

	foo := Foo{Number: 1, Bar: Bar{"title", "name"}}
	fmt.Println(foo)
	foo_marshalled, err := json.Marshal(foo)
	if err == nil {
		fmt.Println(string(foo_marshalled))
	}
}
