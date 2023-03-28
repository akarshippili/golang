package reciver

type Fan struct {
	speed int
}

func (fan *Fan) On() {
	fan.speed = 3
}

func (fan *Fan) Off() {
	fan.speed = 0
}

func (fan *Fan) IncreaseSpeed() {
	fan.speed = (fan.speed + 1) % 5
}

func (fan *Fan) DecreaseSpeed() {
	if fan.speed > 0 {
		fan.speed = (fan.speed - 1)
	}
}
