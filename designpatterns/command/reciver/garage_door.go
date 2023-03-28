package reciver

import "fmt"

type GarageDoor struct{}

func (garageDoor *GarageDoor) Open() {
	fmt.Println("Garage Door Opened")
}

func (garageDoor *GarageDoor) Close() {
	fmt.Println("Garage Door Closed")
}
