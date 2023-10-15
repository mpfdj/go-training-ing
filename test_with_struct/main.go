package main

import "fmt"

type msg struct {
	text string
}

func (m msg) printMessage() {
	fmt.Println(m.text)
}

func main() {
	m := msg{text: "hello"}
	m.printMessage()
}
