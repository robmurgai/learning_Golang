package main

import "fmt"

func thePrinter() {

	//const sample = "\xbd\xb2\x3d\xbc\x20\xe2\x8c\x98"
	const sample = "\xe2\x8c\x98"

	fmt.Printf("\n")
	fmt.Println("Demonstrating different ways to print String: ", sample)
	fmt.Printf("\n")

	fmt.Printf("%%s to display string: %s\n\n", sample)

	fmt.Printf("%x to Display string in Hexademicals(byte code?):\n\n", sample)

	fmt.Println("% x to display string in Hexademicals(byte code?) with sapce between each byte code:")
	fmt.Printf("% x\n\n", sample)

	fmt.Printf("Add \" \" to strings with %q:\n\n", sample)

	fmt.Println("Display charecter code value with %+q:")
	fmt.Printf("%+q\n\n", sample)
}
