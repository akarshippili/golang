package reciver

import "fmt"

type Stereo struct {
	Volume int
}

func (stereo *Stereo) On() {
	stereo.Volume = 3
	fmt.Println("Turned On Stereo")
}

func (stereo *Stereo) Off() {
	stereo.Volume = 0
	fmt.Println("Turned Off Stereo")
}

func (stereo *Stereo) IncreaseVolume() {
	stereo.Volume = (stereo.Volume + 1) % 5
	fmt.Printf("increased volume to %d", stereo.Volume)
}

func (stereo *Stereo) DecreaseVolume() {
	if stereo.Volume > 0 {
		stereo.Volume = (stereo.Volume - 1)
		fmt.Printf("decreased volume to %d", stereo.Volume)
	}
}
