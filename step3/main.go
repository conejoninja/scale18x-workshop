package main

import (
	"image/color"
	"machine"
	"time"

	"tinygo.org/x/drivers/shifter"

	"tinygo.org/x/drivers/ws2812"
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

var colors = [8]color.RGBA{
	color.RGBA{255, 0, 0, 255},
	color.RGBA{255, 255, 0, 255},
	color.RGBA{0, 255, 0, 255},
	color.RGBA{0, 255, 255, 255},
	color.RGBA{0, 0, 255, 255},
	color.RGBA{255, 0, 255, 255},
	color.RGBA{255, 255, 255, 255},
	color.RGBA{0, 0, 0, 255},
}

func main() {
	neo := machine.NEOPIXELS
	neo.Configure(machine.PinConfig{Mode: machine.PinOutput})
	leds := ws2812.New(neo)
	ledColors := make([]color.RGBA, 5)

	buttons := shifter.New(shifter.EIGHT_BITS, machine.BUTTON_LATCH, machine.BUTTON_CLK, machine.BUTTON_OUT)
	buttons.Configure()

	c := 0
	for {
		buttons.Read8Input()
		if buttons.Pins[BUTTON_LEFT].Get() {
			c = 0
		}
		if buttons.Pins[BUTTON_UP].Get() {
			c = 1
		}
		if buttons.Pins[BUTTON_DOWN].Get() {
			c = 2
		}
		if buttons.Pins[BUTTON_RIGHT].Get() {
			c = 3
		}
		if buttons.Pins[BUTTON_SELECT].Get() {
			c = 6
		}
		if buttons.Pins[BUTTON_START].Get() {
			c = 7
		}
		if buttons.Pins[BUTTON_A].Get() {
			c = 4
		}
		if buttons.Pins[BUTTON_B].Get() {
			c = 5
		}

		for i := 0; i < 5; i++ {
			ledColors[i] = colors[c]
		}

		leds.WriteColors(ledColors)
		time.Sleep(time.Millisecond * 30)
	}
}
