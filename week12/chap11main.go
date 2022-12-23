package main

import "gadget"

type Player interface {
	Play(string)
	Stop()
}

func playList(device Player, songs []string) {
	for _, song := range songs {
		device.Play(song)
	}
	device.Stop()
}

func main() {
	player := gadget.TapePlayer{}
	recordPlayer := gadget.TapeRecorder{}
	mcrtape := []string{"인기인기", "인기는 울보", "행복한 인기"}
	mixtape := []string{"투명우산", "좋은날", "너랑나"}
	playList(player, mixtape)
	playList(recordPlayer, mcrtape)
}

/* package main

import "gadget"

func playList(device gadget.TapePlayer, songs []string) {
	for _, song := range songs {
		device.Play(song)
	}
	device.Stop()
}

func main() {
	player := gadget.TapePlayer{}
	recordPlayer := gadget.TapeRecorder{}
	mcrtape := []string{"인기인기", "인기는 울보", "행복한 인기"}
	mixtape := []string{"투명우산", "좋은날", "너랑나"}
	playList(player, mixtape)
	playList(recordPlayer, mcrtape)
}
*/
