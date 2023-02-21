package main

import (
	"machine"
	"time"

	"github.com/zivoy/tinyGoComponents/button"
)

func main() {
	led := machine.LED
	led.Configure(machine.PinConfig{Mode: machine.PinOutput})

	button1 := button.NewToggleButton(button.NewButton(machine.BUTTONA, machine.PinInputPulldown))
	button2 := button.NewButton(machine.BUTTONB, machine.PinInputPulldown)

	combButton := button.NewMultiButton(button.MultiXor, button1, button2)

	for {
		led.Set(combButton.Get())
		time.Sleep(time.Millisecond * 10)
	}
}
