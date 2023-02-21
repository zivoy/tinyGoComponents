package main

import (
	"fmt"
	"image/color"
	"machine"
	"math"
	"time"

	"github.com/zivoy/tinyGoComponents/rotation"
	"github.com/zivoy/tinyGoComponents/utils"

	"tinygo.org/x/drivers/lis3dh"
	"tinygo.org/x/drivers/ws2812"
)

var (
	i2c = machine.I2C1

	neo  = machine.WS2812
	leds [10]color.RGBA

	Reference = rotation.NewReference(utils.Vector3{Z: 1}, utils.Vector3{X: 1})
)

func main() {
	i2c.Configure(machine.I2CConfig{SCL: machine.SCL1_PIN, SDA: machine.SDA1_PIN})

	accel := lis3dh.New(i2c)
	accel.Address = lis3dh.Address1 // address on the Circuit Playground Express
	accel.Configure()
	accel.SetRange(lis3dh.RANGE_2_G)

	println(accel.Connected())

	neo.Configure(machine.PinConfig{Mode: machine.PinOutput})

	ws := ws2812.New(neo)

	for {
		angle, magnitude, _ := Reference.Read(accel)
		// println("X:", int(x*100), "Y:", int(y*100), "Z:", int(z*100))
		for i := range leds {
			leds[i] = color.RGBA{R: 0x04, G: 0x04, B: 0x04}
		}

		println(fmt.Sprintf("angle: %v mag: %v", angle/math.Pi*180, magnitude))

		if magnitude > .9 {
			//angle = angle*-1
			index := int((angle+math.Pi)/(2*math.Pi)*10) % len(leds)
			leds[index] = color.RGBA{R: 0xff, G: 0x00, B: 0x00}
		}

		ws.WriteColors(leds[:])

		time.Sleep(time.Millisecond * 100)
	}
}
