/*
** Write a program that converts strings to Booleans:
**
** The strings “true”, “yes”, or “1” are true.
** The strings “false”, “no”, or “0” are false.
** Display an error message for any other values.
**
 */

package practicefundamentals

import "fmt"

func lesson100(boolString string) {

	var boolValue bool

	switch boolString {
	case "true":
		boolValue = true
		fmt.Printf("Successfully converted String %v to boolean %v\n", boolString, boolValue)
	case "yes":
		boolValue = true
		fmt.Printf("Successfully converted String %v to boolean %v\n", boolString, boolValue)
	case "1":
		boolValue = true
		fmt.Printf("Successfully converted String %v to boolean %v\n", boolString, boolValue)
	case "false":
		boolValue = false
		fmt.Printf("Successfully converted String %v to boolean %v\n", boolString, boolValue)
	case "no":
		boolValue = false
		fmt.Printf("Successfully converted String %v to boolean %v\n", boolString, boolValue)
	case "0":
		boolValue = false
		fmt.Printf("Successfully converted String %v to boolean %v\n", boolString, boolValue)
	default:
		fmt.Printf("Unable to convert %v to booleann\n", boolString)
	}

}
