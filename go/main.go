package main

import "fmt"

type Screen struct {
}

type Phone struct {
	Screen
}

type Reader interface {
	Read() string
}

func (s Screen) Read() string {
	return "Read from screen."
}

func main() {
	var phone Phone
	content := phone.Screen.Read()
	fmt.Printf(content)
}
