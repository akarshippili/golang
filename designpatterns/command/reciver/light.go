package reciver

import "fmt"

type Light struct {
	Brightness int
}

func (light *Light) On() {
	light.Brightness = 10
	fmt.Println("Turned On light")
}

func (light *Light) Off() {
	light.Brightness = 0
	fmt.Println("Turned Off light")
}

func (light *Light) Increase() {
	light.Brightness = (light.Brightness + 1) % 10
}

func (light *Light) Decrease() {
	if light.Brightness > 0 {
		light.Brightness = (light.Brightness - 1)
	}
}
