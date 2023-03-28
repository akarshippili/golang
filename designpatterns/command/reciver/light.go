package reciver

type Light struct {
	brightness int
}

func (light *Light) On() {
	light.brightness = 10
}

func (light *Light) Off() {
	light.brightness = 0
}

func (light *Light) Increase() {
	light.brightness = (light.brightness + 1) % 10
}

func (light *Light) Decrease() {
	if light.brightness > 0 {
		light.brightness = (light.brightness - 1)
	}
}
