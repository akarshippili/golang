package reciver

import "fmt"

type Fan struct {
	Speed int
}

func (fan *Fan) On() {
	fan.Speed = 3
	fmt.Println("Turned On Fan")
}

func (fan *Fan) Off() {
	fan.Speed = 0
	fmt.Println("Turned Off Fan")
}

func (fan *Fan) IncreaseSpeed() {
	if fan.Speed < 5 {
		fan.Speed = fan.Speed + 1
		fmt.Printf("increased speed to %d \n", fan.Speed)
	}
}

func (fan *Fan) DecreaseSpeed() {
	if fan.Speed > 0 {
		fan.Speed = (fan.Speed - 1)
		fmt.Printf("decreased speed to %d \n", fan.Speed)
	}
}
