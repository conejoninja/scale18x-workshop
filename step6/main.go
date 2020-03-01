package main

import (
	"image/color"
	"machine"
	"time"

	"tinygo.org/x/drivers/shifter"
	"tinygo.org/x/tinydraw"

	"tinygo.org/x/drivers/st7735"
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

func main() {
	machine.SPI1.Configure(machine.SPIConfig{
		SCK:       machine.SPI1_SCK_PIN,
		MOSI:      machine.SPI1_MOSI_PIN,
		MISO:      machine.SPI1_MISO_PIN,
		Frequency: 8000000,
	})
	display := st7735.New(machine.SPI1, machine.TFT_RST, machine.TFT_DC, machine.TFT_CS, machine.TFT_LITE)
	display.Configure(st7735.Config{
		Rotation: st7735.ROTATION_90,
	})

	buttons := shifter.New(shifter.EIGHT_BITS, machine.BUTTON_LATCH, machine.BUTTON_CLK, machine.BUTTON_OUT)
	buttons.Configure()

	display.FillScreen(color.RGBA{255, 255, 255, 255})

	circle := color.RGBA{0, 100, 250, 255}
	white := color.RGBA{255, 255, 255, 255}
	ring := color.RGBA{200, 0, 0, 255}
	display.FillScreen(white)
	tinydraw.FilledCircle(&display, 25, 74, 8, circle) // LEFT
	tinydraw.FilledCircle(&display, 55, 74, 8, circle) // RIGHT
	tinydraw.FilledCircle(&display, 40, 59, 8, circle) // UP
	tinydraw.FilledCircle(&display, 40, 89, 8, circle) // DOWN

	tinydraw.FilledCircle(&display, 45, 30, 8, circle)  // SELECT
	tinydraw.FilledCircle(&display, 120, 30, 8, circle) // START

	tinydraw.FilledCircle(&display, 120, 80, 8, circle) // B
	tinydraw.FilledCircle(&display, 135, 65, 8, circle) // A

	for {
		buttons.Read8Input()
		if buttons.Pins[BUTTON_SELECT].Get() {
			tinydraw.Circle(&display, 45, 30, 10, ring)
		} else {
			tinydraw.Circle(&display, 45, 30, 10, white)
		}
		if buttons.Pins[BUTTON_START].Get() {
			tinydraw.Circle(&display, 120, 30, 10, ring)
		} else {
			tinydraw.Circle(&display, 120, 30, 10, white)
		}
		if buttons.Pins[BUTTON_A].Get() {
			tinydraw.Circle(&display, 135, 65, 10, ring)
		} else {
			tinydraw.Circle(&display, 135, 65, 10, white)
		}
		if buttons.Pins[BUTTON_B].Get() {
			tinydraw.Circle(&display, 120, 80, 10, ring)
		} else {
			tinydraw.Circle(&display, 120, 80, 10, white)
		}
		if buttons.Pins[BUTTON_LEFT].Get() {
			tinydraw.Circle(&display, 25, 74, 10, ring)
		} else {
			tinydraw.Circle(&display, 25, 74, 10, white)
		}
		if buttons.Pins[BUTTON_RIGHT].Get() {
			tinydraw.Circle(&display, 55, 74, 10, ring)
		} else {
			tinydraw.Circle(&display, 55, 74, 10, white)
		}
		if buttons.Pins[BUTTON_UP].Get() {
			tinydraw.Circle(&display, 40, 59, 10, ring)
		} else {
			tinydraw.Circle(&display, 40, 59, 10, white)
		}
		if buttons.Pins[BUTTON_DOWN].Get() {
			tinydraw.Circle(&display, 40, 89, 10, ring)
		} else {
			tinydraw.Circle(&display, 40, 89, 10, white)
		}
		time.Sleep(50 * time.Millisecond)
	}
}
