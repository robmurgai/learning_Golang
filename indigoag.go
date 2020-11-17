// The challenge is to create a program that computes some basic
// statistics on a collection of small positive integers. You can
// assume all values will be less than 1,000.
package main

// You cannot import a library that solves it instantly
import (
	//For Printing to the terminal
	"fmt"
)

// func main() {

// 	fmt.Printf("######### START ###########\n\n")

// 	capture := dataCapture()
// 	fmt.Println("Debug Main(): Initalize capture - An indigoValue object")
// 	capture.prettyPrint()

// 	capture.add(3)
// 	capture.add(9)
// 	capture.add(3)
// 	capture.add(4)
// 	capture.add(6)
// 	// capture.add(0)
// 	// capture.add(500)
// 	// capture.add(1000)
// 	// capture.add(1001)
// 	// capture.add(1002)
// 	// capture.add(-1)

// 	fmt.Println("Debug Main(): Add 3, 9, 3, 4, 6")
// 	capture.prettyPrint()

// 	stats := capture.buildStats()
// 	fmt.Println("Debug main(): Initialize stats - An indigoStats Object")
// 	// fmt.Printf("%#v", stats)
// 	stats.prettyPrint()

// 	//stats.less(4) should return 2 (only two values 3, 3 are less than 4)
// 	fmt.Printf("stats.less(4): %v\n", stats.less(4))

// 	// stats.between(3, 6) # should return 4 (3, 3, 4 and 6 are between 3 and 6)
// 	fmt.Printf("stats.between(3, 6): %v\n", stats.between(3, 6))

// 	//# stats.less(4) should return 2 (6 and 9 are the only two values greater than 4)
// 	fmt.Printf("stats.greater(4): %v\n", stats.greater(4))

// 	fmt.Printf("\n\n########## END ############\n")
// }

// indigoValue contains a slice of integers and the total number of integers added.
type indigoValue struct {
	idov []int //Slice of integers
	idol int   //Number of integers added to this inidgoValue object
}

// indigoValue contains a two dimentional slice of integers and the total number of integers added.
type indigoStats struct {
	isv [][]int // the second will have two values: the number of occurences of index, and the ranking of the index
	isl int     // the number of elements in the struct
}

// DataCapture accepts initializes an indigoValue struct with a slice of 1001 elements.
// Instead of each element referring to a numberical_value added to the struct,
// the index of each element refers to the numberical value added to the struct and the
// value at that index refers to the number of occurences of the said numerical_value.
func dataCapture() indigoValue {

	//Initialize a temporary slice with a lenght of 1001, so the index goes from 0 to 1000.  By default all value are 0
	tmp := make([]int, 1001)

	//Add the temp array to result, with a lenght 0.
	result := indigoValue{idov: tmp, idol: 0}

	return result
}

// Add increments the value of every index in ido.idov, where the index is the number being passed.
// Here the number being passed is used as an index and the value of that index is incremented.
// Hence if the number passed is 3, we essentially are doing idov[3] = idov[3]++
// - The methods add(), less(), greater(), and between() should have constant time O(1)
func (ido *indigoValue) add(numericalValue int) {

	if numericalValue < 1001 && numericalValue >= 0 {
		ido.idov[numericalValue] = ido.idov[numericalValue] + 1
		ido.idol++
	} else {
		fmt.Printf("add(): %v is outsid the 0 to 1000 range and will not be added\n", numericalValue)
	}
}

// buildStats returns an object for
// querying statistics about the inputs. Specifically, the returned object supports
// querying how many numbers in the collection are less than a value,
// greater than a value, or within a range.
func (ido *indigoValue) buildStats() indigoStats {

	// Initialize a slice of type integer and length same as the number of values in object ido
	tmp := make([][]int, 1001)
	var stats = indigoStats{isv: tmp, isl: 0}

	//The for loop will go over all 1001 values
	for i, val := range ido.idov {
		if val != 0 {
			stats.isv[i] = []int{val, stats.isl}
			stats.isl = stats.isl + val
		}
	}

	//fmt.Printf("Debug buildStats(): stats: %+v\n", stats)
	return stats
}

// stats.less(4) # should return 2 (only two values 3, 3 are less than 4)
// - The methods add(), less(), greater(), and between() should have constant time O(1)
func (stats *indigoStats) less(value int) int {

	// assume stats has a sorted list isv where the value is the stat index of the key.
	// in that case the value of the particular element could be its position in the list.
	var result = stats.isv[value][1]
	return result
}

// stats.between(3, 6) # should return 4 (3, 3, 4 and 6 are between 3 and 6)
// - The methods add(), less(), greater(), and between() should have constant time O(1)
func (stats *indigoStats) between(start int, end int) int {
	// assume stats has a sorted list isv where the value is the stat index of the key.
	// in that case the value of the particular element could be its position in the list.
	var result = stats.isv[end][1] - stats.isv[start][1] + 1
	return result
}

// stats.greater(4) # should return 2 (6 and 9 are the only two values greater than 4)
// - The methods add(), less(), greater(), and between() should have constant time O(1)
func (stats *indigoStats) greater(value int) int {
	// assume stats has a sorted list isv where the value is the stat index of the key.
	// in that case the value of the particular element could be its position in the list.
	var result = stats.isl - stats.isv[value][1] - 1
	return result
}

func (ido *indigoValue) prettyPrint() {
	result := fmt.Sprintf("indigoValue Object has %v elements\n", ido.idol)
	if ido.idol > 0 {
		result += fmt.Sprintf("%-9v%-12v\n", "Number", "Occurance")
	}
	for i, val := range ido.idov {
		if val != 0 {
			result += fmt.Sprintf("%-9v%-12v\n", i, val)
		}
	}
	result += "\n"
	fmt.Printf(result)
}

func (stats *indigoStats) prettyPrint() {
	result := fmt.Sprintf("indigoStats Object has %v elements\n", stats.isl)
	if stats.isl > 0 {
		result += fmt.Sprintf("%-9v%-12v%v\n", "Number", "Occurance", "At Index")
	}
	for i, val := range stats.isv {
		if val != nil {
			result += fmt.Sprintf("%-9v%-12v%v\n", i, val[0], val[1])
		}
	}
	result += "\n"
	fmt.Printf(result)
}
