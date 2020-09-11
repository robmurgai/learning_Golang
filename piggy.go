/**
 ** Experiment: piggy.go
 ** Write a new piggy bank program that uses integers to track the number of cents rather than dollars. Randomly place nickels (5¢), dimes (10¢), and
 ** quarters (25¢) into an empty piggy bank until it contains at least $20.
 **
 ** Display the running balance of the piggy bank after each deposit in dollars (for example, $1.05).
 **
 ** Tip: If you need to find the remainder of dividing two numbers, use modulus (%).
 **/

package main

import (
	"fmt"
	"math/rand"
)

func lesson700() {

	const maxBalance = 20
	balanceDollars := 0
	balanceCents := 0
	fmt.Printf("Balance: $%v.%02v\n\n", balanceDollars, balanceCents)

	for balanceDollars < maxBalance {
		additionalAmount := 0
		switch rand.Intn(3) {
		case 0:
			additionalAmount = 5
		case 1:
			additionalAmount = 10
		case 2:
			additionalAmount = 25
		}

		fmt.Printf("Debug: Adding %v cents\n", additionalAmount)
		balanceCents += additionalAmount
		balanceDollars += balanceCents / 100
		balanceCents = balanceCents % 100
		fmt.Printf("Balance: $%v.%02v\n\n", balanceDollars, balanceCents)
	}
}
