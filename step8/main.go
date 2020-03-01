package main

import (
	"machine"
	"time"

	"tinygo.org/x/drivers/shifter"
)

const (
	BUTTON_LEFT = iota
	BUTTON_UP
	BUTTON_DOWN
	BUTTON_RIGHT
	BUTTON_SELECT
	BUTTON_START
	BUTTON_A
	BUTTON_B
)

var bzrPin machine.Pin

func main() {
	speaker := machine.SPEAKER_ENABLE
	speaker.Configure(machine.PinConfig{Mode: machine.PinOutput})
	speaker.High()

	bzrPin = machine.A0
	bzrPin.Configure(machine.PinConfig{Mode: machine.PinOutput})

	buttons := shifter.New(shifter.EIGHT_BITS, machine.BUTTON_LATCH, machine.BUTTON_CLK, machine.BUTTON_OUT)
	buttons.Configure()

	for {
		buttons.Read8Input()
		if buttons.Pins[BUTTON_LEFT].Get() {
			tone(329)
		}
		if buttons.Pins[BUTTON_UP].Get() {
			tone(369)
		}
		if buttons.Pins[BUTTON_DOWN].Get() {
			tone(523)
		}
		if buttons.Pins[BUTTON_RIGHT].Get() {
			tone(739)
		}
		if buttons.Pins[BUTTON_A].Get() {
			tone(1046)
		}
		if buttons.Pins[BUTTON_B].Get() {
			tone(1975)
		}
		if buttons.Pins[BUTTON_SELECT].Get() {
			tone(2637)
		}
		if buttons.Pins[BUTTON_START].Get() {
			tone(5274)
		}
	}
}

func tone(tone int) {
	for i := 0; i < 10; i++ {
		bzrPin.High()
		time.Sleep(time.Duration(tone) * time.Microsecond)

		bzrPin.Low()
		time.Sleep(time.Duration(tone) * time.Microsecond)
	}
}
