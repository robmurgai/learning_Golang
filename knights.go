// Practice: Nil
// A knight blocks Arthur’s path. Our hero is empty-handed, represented by a nil value for leftHand *item. Implement a character struct
// with methods such as pickup(i *item) and give(to *character) and write a script that has Arthur pick up an item and give it to the knight,
// displaying an appropriate description for each action
//
// func main() {
//
// 	fmt.Printf("########## Starting Exercise ##########\n\n")
//
// 	var arthur, lancealot charector
// 	arthur.name = "Arthur"
// 	lancealot.name = "Lancealot"
//
// 	var excalibur *item
// 	excalibur = &item{itemName: "Excalibur"}
//
// 	fmt.Printf("%v runs into %v in the forest\n", arthur.name, lancealot.name)
// 	//If Arthur is empty handed, have him pick up the Excalibur
// 	if arthur.leftHand == nil {
// 		fmt.Printf("%v is empty handed\n", arthur.name)
// 		arthur.pickup(excalibur)
// 	}
// 	fmt.Printf("%v has %v\n", arthur.name, arthur.leftHand.itemName)
//
// 	//Have Aurthur give Excalibur to Lancelot.
// 	arthur.give(&lancealot)
//
// 	fmt.Printf("%v now has %v\n", lancealot.name, lancealot.leftHand.itemName)
// 	fmt.Printf("%v has %v\n", arthur.name, arthur.leftHand)
//
// 	fmt.Printf("\n########## Ending Exercise ##########\n")
//
// }

package main

import "fmt"

type item struct {
	itemName string
}

type charector struct {
	leftHand *item
	name     string
}

func (c *charector) pickup(i *item) {

	// fmt.Println("\n\nWhile not needed for the exercise, the following print statements are to help understand pointers")
	// fmt.Printf("\ni: \n")
	// fmt.Printf("  i is a pointer of type %T and is currently = %+v\n", i, i)
	// fmt.Printf("  *i is the value i points to, the value is of type %T and currently = %+v\n\n", *i, *i)

	// fmt.Println("itemName")
	// fmt.Printf("  (*i).itemName is of type %T and contains value %v\n", (*i).itemName, (*i).itemName)
	// fmt.Printf("  i.itemName is of type %T and contains value %v\n\n", i.itemName, i.itemName)

	// fmt.Println("Summary: ")
	// fmt.Printf("  As you can see i.itemName and (*i).itemname are returning the same type: %T and Value: %v\n", i.itemName, i.itemName)
	// fmt.Println("  Because Go Compiler trqnslates (*i).itemname to i.itemName to enable readability of code.")
	// fmt.Printf("  Hence i is a pointer of type %T, it points to a specific item struct with itemName = %v\n\n\n", i, i.itemName)

	c.leftHand = i
	fmt.Printf("%v picks up %v\n", c.name, c.leftHand.itemName)

}

func (c *charector) give(to *charector) {
	to.leftHand = c.leftHand
	c.leftHand = nil
	fmt.Printf("%v gives %v to %v\n", c.name, to.leftHand.itemName, to.name)
}
