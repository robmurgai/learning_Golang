package main

import (
	"fmt"
	"math/rand"
	"time"
)

//  Use this in main.go as the funtion to understand what calliberate is doing
//  func main() {
// 	fmt.Printf("######### START ###########\n\n")

// 	var newOffset kelvin = 10

// 	sensor := calibrate(realSensor, newOffset)
// 	fmt.Printf("newOffset(%v) & Sensor reading is: ", newOffset)
// 	fmt.Println(sensor())

// 	newOffset = kelvin(15)
// 	fmt.Printf("newOffset(%v) & Sensor reading is: ", newOffset)
// 	fmt.Println(sensor())

// 	fmt.Printf("newOffset(%v) & Sensor reading is: ", newOffset)
// 	fmt.Println(calibrate(realSensor, newOffset))

// 	fmt.Printf("\n\n########## END ############\n")
//  }

// sensor function type
type sensor func() kelvin

func realSensor() kelvin {
	return 0
}

func calibrate(s sensor, offset kelvin) sensor {
	return func() kelvin {
		return s() + offset
	}
}

func measureTemperature(samples int, sensor func() kelvin) {
	for i := 0; i < samples; i++ {
		k := sensor()
		fmt.Printf("%v° K\n", k)
		time.Sleep(time.Second)
	}
}

func fakeSensor() kelvin {
	return kelvin(rand.Intn(151) + 150)
}
