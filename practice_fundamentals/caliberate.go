// Pracitice Passing funcations as parameters.
// In Go you can assign functions to variables, pass functions to functions, and write functions that return
// functions. This ability is defined as Functions being first-class, as in they work in all the places that integers, strings, and other
// types work.
// Given the following code:
// 		package practicefundamentals
//
// 		import "fmt"
//
// 		type kelvin float64
//
// 		// sensor function type
// 		type sensor func() kelvin
//
// 		func realSensor() kelvin {
//     		return 0
// 		}
//
// 		func calibrate(s sensor, offset kelvin) sensor {
//     		return func() kelvin {
//          		return s() + offset
//     		}
// 		}
//
// 		func main() {
//     		sensor := calibrate(realSensor, 5)
//     		fmt.Println(sensor())
// 		}
//
// Exercise 1) Rather than passing 5 as an argument to calibrate, declare and pass a variable. Modify the variable and you’ll notice that
// calls to sensor() still result in 5. That’s because the offset parameter is a copy of the argument (pass by value).
//
// Exercise 2) Use calibrate with the following fakeSensor function to create a new sensor function. Call the new sensor function multiple
// times and notice that the original fakeSensor is still being called each time, resulting in random values.
//
// 		func fakeSensor() kelvin {
//     		return kelvin(rand.Intn(151) + 150)
// 		}

package practicefundamentals

import (
	"fmt"
	"math/rand"
	"time"
)

//  Use this in main.go as the function to understand what calliberate is doing
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
