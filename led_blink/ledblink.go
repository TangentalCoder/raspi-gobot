package main

import (
	"fmt"
	"time"

	"github.com/stianeikeland/go-rpio"
)

func main() {
	fmt.Println("opening gpio")
	err := rpio.Open()
	if err != nil {
		panic(fmt.Sprint("unable to open gpio", err.Error()))
	}

	defer rpio.Close()

	morse(18)

}
func morse(pinNumber int) {
	pin := rpio.Pin(18)
	pin.Output()

	// for x := 0; x < 20; x++ {
	// 	pin.Toggle()
	// 	time.Sleep(time.Second / 5)
	// }
	//showShort(pin)
	//pause(pin)
	//showLong(pin)
	// H
	showShort(pin)
	pause(pin)
	showShort(pin)
	pause(pin)
	showShort(pin)
	pause(pin)
	showShort(pin)
	letterpause(pin)

	// I
	showShort(pin)
	pause(pin)
	showShort(pin)
	pause(pin)
}

func pause(pin rpio.Pin) {
	pin.Low()
	time.Sleep(time.Second / 4)
}

func letterpause(pin rpio.Pin) {
	pin.Low()
	time.Sleep(time.Second / 2)
}

func wordpause(pin rpio.Pin) {
	pin.Low()
	time.Sleep(time.Second)
}

func showShort(pin rpio.Pin) {
	pin.Low()
	pin.High()
	time.Sleep(time.Second / 5)
	pin.Low()
}

func showLong(pin rpio.Pin) {
	pin.Low()
	pin.High()
	time.Sleep(time.Second)
	pin.Low()
}
