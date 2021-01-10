// Practice concurrency in Go by calling a worker node and using time.After() and time.Sleep funcations.
// Using the following code as a starting point, change the delay time so it gets a half a second longer with each move.
//
// func worker() {
//     pos := image.Point{X: 10, Y: 10}                   1
//     direction := image.Point{X: 1, Y: 0}               2
//     next := time.After(time.Second)
//     for {
//         select {
//         case <-next:
//             pos = pos.Add(direction)
//             fmt.Println("current position is ", pos)   3
//             next = time.After(time.Second)
//         }
//     }
// }
//
// Use the following calls in main for this exercise
// func main() {

// 	fmt.Printf("########## Starting Exercise ##########\n\n")

// 	go worker()
// 	time.Sleep(5 * time.Second)

// 	fmt.Printf("\n########## Ending Exercise ##########\n")

// }

package practicefundamentals

import (
	"fmt"
	"image"
	"time"
)

func worker() {
	fmt.Println("Starting workder routine")
	pos := image.Point{X: 10, Y: 10}
	direction := image.Point{X: 1, Y: 0}
	second := time.Second
	fmt.Println("Will come back in ", second, " seconds")
	next := time.After(second)
	for {
		select {
		case <-next:
			pos = pos.Add(direction)
			fmt.Println("current position is ", pos)
			second += ((time.Second) / 2)
			fmt.Println("Will come back in ", second, " seconds")
			next = time.After(second)
		}
	}
}
