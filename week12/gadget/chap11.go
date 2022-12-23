package gadget

import "fmt"

type TapePlayer struct {
	Batteries string
}

type TapeRecorder struct {
	Microphones int
}

func (t TapePlayer) Play(song string) {
	fmt.Println("재생중", song)
}

func (t TapePlayer) Stop() {
	fmt.Println("정지됨!")
}

func (t TapeRecorder) Play(song string) {
	fmt.Println("녹음중", song)
}

func (t TapeRecorder) Stop() {
	fmt.Println("정지됨!")
}
