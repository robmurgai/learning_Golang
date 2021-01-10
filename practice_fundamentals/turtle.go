// Experiment: turtle.go
// Write a program with a turtle that can move up, down, left, or right. The turtle should store an (x, y) location where positive values
// go down and to the right. Use methods to increment/decrement the appropriate variable. A main function should exercise the methods
// you’ve written and print the final location.

// Tip
// Method receivers will need to use pointers in order to manipulate the x and y values.

package practicefundamentals

//
// func main() {
//
// 	t := turtle{x: 0, y: 0}
// 	fmt.Printf("Turtle t is at: %+v\n", t)
//
// 	for i := 0; i < 10; i++ {
// 		switch rand.Intn(4) {
// 		case 0:
// 			fmt.Println("Moving Up")
// 			t.up()
// 		case 1:
// 			fmt.Println("Moving Down")
// 			t.down()
// 		case 2:
// 			fmt.Println("Moving Right")
// 			t.right()
// 		case 3:
// 			fmt.Println("Moving Left")
// 			t.left()
// 		}
// 	}
//
// 	fmt.Printf("Turtle t is at: %+v\n", t)

// The turtle should store an (x, y) location where positive values go down and to the right.
type turtle struct {
	x int
	y int
}

func (t *turtle) down() {
	t.y++
}
func (t *turtle) up() {
	t.y--
}

func (t *turtle) right() {
	t.x++
}

func (t *turtle) left() {
	t.x--
}
