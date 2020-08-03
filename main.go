package main

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
	"unicode/utf8"
)

func main() {
	fmt.Printf("######### START ###########\n\n")

	var newOffset kelvin = 10

	sensor := calibrate(realSensor, newOffset)
	fmt.Printf("newOffset(%v) & Sensor reading is: ", newOffset)
	fmt.Println(sensor())

	newOffset = kelvin(15)
	fmt.Printf("newOffset(%v) & Sensor reading is: ", newOffset)
	fmt.Println(sensor())

	fmt.Printf("newOffset(%v) & Sensor reading is: ", newOffset)
	fmt.Println(calibrate(realSensor, newOffset))

	measureTemperature(3, fakeSensor)

	fmt.Printf("\n\n########## END ############\n")
}

/**
** To send ciphered messages, write a program that ciphers plain text using a keyword:
**
** plainText := "your message goes here"
** keyword := "GOLANG"
** Bonus: rather than write your plain text message in uppercase letters with no spaces, use the strings.Replace and strings.ToUpper functions to remove   ** spaces and uppercase the string before you cipher it.
**
** Once you’ve ciphered a plain text message, check your work by deciphering the ciphered text with the same keyword.
**/

func lesson112(plainText string, cipherKeyword string) {

	var cipheredText string
	var cipherKeywordSize = utf8.RuneCountInString(cipherKeyword)
	var cipherKeywordIndex int

	fmt.Printf("plainText: %v of size %v\n", plainText, len(plainText))
	fmt.Printf("cipherKeyword: %v of size %v\n", cipherKeyword, cipherKeywordSize)

	plainText = strings.Replace(plainText, " ", "", -1)
	fmt.Printf("Remiving Spaces - plainText: %v\n", plainText)

	plainText = strings.ToUpper(plainText)
	fmt.Printf("Translating to Upper Case - plainText: %v of size %v\n", plainText, len(plainText))

	for i, plainTextAtRune := range plainText {
		fmt.Printf("\nIndex: %v\n", i)
		fmt.Printf("plainTextAtRune[%v]: %v\n", i, string(plainTextAtRune))

		cipherKeywordIndex = i % cipherKeywordSize
		fmt.Printf("cipherKeywordAtRune[%v]: %v\n", cipherKeywordIndex, string(cipherKeyword[cipherKeywordIndex]))

		a := plainTextAtRune - 'A'
		b := cipherKeyword[cipherKeywordIndex] - 'A'
		c := a + rune(b)
		d := c % 26
		e := d + 'A'
		cipheredTextRune := string(e)

		fmt.Printf("plainTextAtRune Value relative to 'A': %v\n", a)
		fmt.Printf("cipherKeywordAtRune Value relative to 'A': %v\n", b)
		fmt.Printf("Adding them both to create Cipher Rune Value: %v\n", c)
		fmt.Printf("Modulo 26 incase the value is bigger than 26 (26 chars in the Alphabet): %v\n", d)
		fmt.Printf("Cipher Rune Value relative to 'A': %v\n", e)
		fmt.Printf("Converting Rune Value to String - cipheredTextAtRune[%v]: %v\n", i, cipheredTextRune)

		cipheredText += string(cipheredTextRune)

	}

	fmt.Printf("\ncipheredText: %v\n\n", cipheredText)
}

/*
** EXPERIMENT: DECIPHER.GO
** Write a program to decipher the ciphered text shown in table 11.2. To keep it simple, all characters are uppercase English letters for both the text and
** keyword:
**
** 	cipherText := "CSOITEUIWUIZNSROCNKFD"   //CSOITE UIWUIZ NSROCN KFD
** 	keyword := "GOLANG"
**
** The strings.Repeat function may come in handy. Give it a try, but also complete this exercise without importing any packages other than fmt to print the
** deciphered message.  Try this exercise using range in a loop and again without it. Remember that the range keyword splits a string into runes, whereas an
** index like keyword[0] results in a byte.   To wrap around at the edges of the alphabet, the Caesar cipher exercise made use of a comparison. Solve this
** exercise without any if statements by using modulus (%).
**
** Tip
**  - You can only perform operations on values of the same type, but you can convert one type to the other (string, byte, rune).
**  - If you recall, modulus gives the remainder of dividing two numbers. For example, 27 % 26 is 1, keeping numbers within the 0–25 range. Be careful with
**  negative numbers, though, as -3 % 26 is still -3.
**
 */

func lesson110(cipherText string, cipherKeyword string) {

	var modOper = len(cipherKeyword)
	var plainText string

	fmt.Printf("The cipherText is: %v and keyword is: %v\n", cipherText, cipherKeyword)
	//fmt.Printf("Byte code value of 'A':  %v\n\n", byte('A'))

	for i := 0; i < len(cipherText); i++ {

		keywordNum := i % modOper
		cipherTextByte := cipherText[i]
		cipherKeywordByte := cipherKeyword[keywordNum]
		plainTextByte := (cipherTextByte-cipherKeywordByte+26)%26 + 'A'
		plainTextChar := string(plainTextByte)
		plainText += plainTextChar

		// fmt.Printf("cipherText[%v]: %c with byte value of %v\n", i, cipherText[i], cipherTextByte)
		// fmt.Printf("cipherKeword[%v]: %c with byte value of %v\n", keywordNum, cipherKeyword[keywordNum], cipherKeywordByte)
		// fmt.Printf("decipherText[%v]: %v with byte value of %v\n\n", i, plainTextChar, plainTextByte)
	}

	fmt.Printf("The decipherText is: %v of size %v\n", plainText, len(plainText))

}

/*
** EXPERIMENT: DECIPHER.GO
** Write a program to decipher the ciphered text shown in table 11.2. To keep it simple, all characters are uppercase English letters for both the text and
** keyword:
**
** 	cipherText := "CSOITEUIWUIZNSROCNKFD"   //CSOITE UIWUIZ NSROCN KFD
** 	keyword := "GOLANG"
**
** The strings.Repeat function may come in handy. Give it a try, but also complete this exercise without importing any packages other than fmt to print the
** deciphered message.  Try this exercise using range in a loop and again without it. Remember that the range keyword splits a string into runes, whereas an
** index like keyword[0] results in a byte.   To wrap around at the edges of the alphabet, the Caesar cipher exercise made use of a comparison. Solve this
** exercise without any if statements by using modulus (%).
**
** Tip
**  - You can only perform operations on values of the same type, but you can convert one type to the other (string, byte, rune).
**  - If you recall, modulus gives the remainder of dividing two numbers. For example, 27 % 26 is 1, keeping numbers within the 0–25 range. Be careful with
**  negative numbers, though, as -3 % 26 is still -3.
**
 */

func lesson111(cipherText string, cipherKeyword string) {

	var modOper = len(cipherKeyword)
	var plainText string

	fmt.Printf("The cipherText is: %v and keyword is: %v\n", cipherText, cipherKeyword)
	//fmt.Printf("Byte code value of 'A':  %v\n\n", byte('A'))

	for i, cipherTextRune := range cipherText {

		keywordNum := i % modOper
		cipherKeywordAtRune := rune(cipherKeyword[keywordNum])
		plainTextAtRune := (cipherTextRune-cipherKeywordAtRune+26)%26 + 'A'
		plainTextChar := string(plainTextAtRune)
		plainText += plainTextChar

		// fmt.Printf("cipherText[%v]: %c with rune value of %v\n", i, cipherText[i], cipherTextRune)
		// fmt.Printf("cipherKeword[%v]: %c with rune value of %v\n", keywordNum, cipherKeyword[keywordNum], cipherKeywordAtRune)
		// fmt.Printf("decipherText[%v]: %v with rune value of %v\n\n", i, plainTextChar, plainTextAtRune)

	}

	fmt.Printf("The decipherText is: %v of size %v\n", plainText, len(plainText))

}

/*
** Write a program that converts strings to Booleans:
**
** The strings “true”, “yes”, or “1” are true.
** The strings “false”, “no”, or “0” are false.
** Display an error message for any other values.
**
 */
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

/**
** What are runes? rune is what we store the unicode value of each charter known to us today.  Runes are int32, just like Bytes are uint8.
** ASCII which is a subset of Unicode set is stored in uint8 aka Bytes, but because we don't have enough bytes to store millions of charecters from
** languages and non language charecter sets (think smileyface charecter) runes are used to store the chars are they use int32
**
** Strings in Go are encoded with UTF-8, one of several encodings for Unicode code points. UTF-8 is an efficient variable length encoding where a single  ** code point may use 8 bits, 16 bits, or 32 bits. By using a variable length encoding, UTF-8 makes the transition from ASCII straightforward, because
** ASCII characters are identical to their UTF-8 encoded counterparts
**
** Strings use a variable length encoding called UTF-8, where each character consumes 1–4 bytes.
**
**/

func lesson992() {

	var AasRune rune = 'A'
	var AasChar = 'A'
	var AasLiteralString = "A"

	fmt.Printf("var AasRune rune = 'A'\n")
	fmt.Printf("AasRune: %v\n\n", AasRune)

	fmt.Printf("var AasChar = 'A'\n")
	fmt.Printf("AasChar: %v\n\n", AasChar)

	fmt.Printf("var AasLiteralString = \"A\"\n")
	fmt.Printf("AasLiteralString: %v\n\n", AasLiteralString)

	fmt.Printf("You can print the unicode value of anything on your keyboard with %v \n", "%v")
	fmt.Printf("Like the unicode value of '*': %v\n\n", '*')

	c := 'a'
	fmt.Printf("Char c: %c\n", c)

	c = c + 3
	fmt.Printf("c = c + 3\n")
	fmt.Printf("Char c: %c\n", c)

}

/*
** Cipher the Spanish message “Hola Estación Espacial Internacional” with ROT13. Modify the following listing to use the range keyword.
**
** for i := 0; i < len(message); i++ {        1
**     c := message[i]
**     if c >= 'a' && c <= 'z' {              2
**         c = c + 13
**         if c > 'z' {
**             c = c - 26
**         }
**     }
**     fmt.Printf("%c", c)
** }
** Now when you use ROT13 on Spanish text, characters with accents are preserved.
 */

func lesson991() {

	message_original := "Hola Estación Especial Internacional"
	var message_new string
	var message_new_with_range string

	// fmt.Printf("Tip: ASCII String of %c is %[1]v\n", 'o')
	// fmt.Printf("Tip: ASCII String of %c is %[1]v\n", 'ó')
	// fmt.Printf("Tip: ASCII String of %c is %[1]v\n", 'a')
	// fmt.Printf("Tip: ASCII String of %c is %[1]v\n", 'A')
	// fmt.Printf("Tip: ASCII String of %c is %[1]v\n", 'z')
	// fmt.Printf("Tip: ASCII String of %c is %[1]v\n", 'Z')

	fmt.Printf("\nThe original Message was:       %v\n", message_original)

	//fmt.Printf("\n\n (##) O NR\n")

	for _, c := range message_original {
		if ('a' <= c && c <= 'm') || ('A' <= c && c <= 'M') {
			//fmt.Printf(" (%2v) %c", i, c)
			c += 13
			//fmt.Printf(" %c ASCII String %[1]v\n ", c)
		} else if ('n' <= c && c <= 'z') || ('N' <= c && c <= 'Z') {
			//fmt.Printf("(%2v) %c", i, c)
			c -= 13
			//fmt.Printf(" %c ASCII String %[1]v\n ", c)
		} else {
			//fmt.Printf("(%2v) %c", i, c)
			//fmt.Printf(" %c Unchanged ASCII String %[1]v\n ", c)

		}
		//fmt.Printf("%v %c\n", i, c)
		message_new_with_range += string(c)
	}

	//fmt.Printf("Processing the For Loop:\n\n")
	//fmt.Printf("(##) O N\n")

	for i := 0; i < len(message_original); i++ {
		c := message_original[i]
		if ('a' <= c && c <= 'm') || ('A' <= c && c <= 'M') {
			//fmt.Printf("(%2v) %c", i, c)
			c += 13
			//fmt.Printf(" %c\n", c)
		} else if ('n' <= c && c <= 'z') || ('N' <= c && c <= 'Z') {
			//fmt.Printf("(%2v) %c", i, c)
			c -= 13
			//fmt.Printf(" %c\n", c)
		} else {
			//fmt.Printf("(%2v) %c", i, c)
			//fmt.Printf(" %c Unchanged\n ", c)

		}
		message_new += string(c)
		//mt.Printf("%c", c)
	}

	for i := 0; i < len(message_original); i++ {
		//fmt.Printf("\n(%v) Looks like the charector %c in original message", i, message_original[i])
		//fmt.Printf("\nIs being turned to %c in new message   ", message_new[i])
		//fmt.Printf("\nIs being turned to %c in Range message \n", message_new_with_range[i])
	}

	fmt.Printf("The new with ranges message is: %v\n", message_new_with_range)
	fmt.Printf("The new message is:             %v\n", message_new)

}

/*
** Decipher the quote from Julius Caesar:
**
** L fdph, L vdz, L frqtxhuhg.
**
** Julius Caesar
**
** Your program will need to shift uppercase and lowercase letters by –3. Remember that 'a' becomes 'x', 'b' becomes 'y', and 'c' becomes 'z',
** and likewise for uppercase letters.
 */

func lesson990() {
	message_original := "L fdph, L vdz, L frqtxhuhg."
	var message_new string

	fmt.Printf("\nThe original Message was:  %v\n", message_original)
	//fmt.Printf("Processing the For Loop:\n\n")

	for i := 0; i < len(message_original); i++ {
		c := message_original[i]
		if ('d' <= c && c <= 'z') || ('D' <= c && c <= 'Z') {
			//fmt.Printf("(%v)Looks like the charector %c", i, c)
			c -= 3
			//fmt.Printf(" is being turned to %c\n", c)
		} else if ('a' <= c && c <= 'c') || ('A' <= c && c <= 'C') {
			//fmt.Printf("(%v)Oh Full circle, Looks like the charector %c", i, c)
			c += 23
			//fmt.Printf(" is being turned to %c\n", c)
		} else {
			//fmt.Printf("(%v)Wait!!, %c wasn't between 'a' 'A' or 'z' 'Z'.", i, c)
			//fmt.Printf(" Leaving at as %c\n", c)

		}
		message_new += string(c)
		//fmt.Printf("%c", c)
	}

	fmt.Printf("The new message is:        %v\n", message_new)

}

/*
** Print the letters in the strong "¿Cómo estás?" using runes and bytes.
** For Bytes, use the len function to determine the length for a variety of types. In this case, len returns the length of a string in bytes.

** For runes, Go language provides the range keyword to iterate over a variety of collections and it can also decode UTF-8 encoded strings,
** as shown in the following listing.

** for i, c := range question {
**     fmt.Printf("%v %c\n", i, c)
** }

** On each iteration, the variables i and c are assigned to an index into the string and the code point (rune) at that position.
** If you don’t need the index, the blank identifier (an underscore) allows you to ignore it.
 */

func lesson980() {

	question := "¿Cómo estás?"

	fmt.Printf("Bytes (8 bit):  ")
	for i := 0; i < len(question); i++ {
		c := question[i]
		fmt.Printf("%c", c)
	}

	fmt.Printf("\nRune (32bit):  ")
	for _, c := range question {
		fmt.Printf("%c ", c)
	}

}

/*
** The first step to supporting other languages is to decode characters to the rune type before manipulating them. Fortunately, Go has functions and
** language features for decoding UTF-8 encoded strings.
**
** The utf8 package provides functions to determine the length of a string in runes rather than bytes and to decode the first character of a string.
** The DecodeRuneInString function returns the first character and the number of bytes the character consumed.
 */

func lesson970() {

	question := "¿Cómo estás?"

	fmt.Println(len(question), "bytes")
	fmt.Println(utf8.RuneCountInString(question), "runes")

	c, size := utf8.DecodeRuneInString(question)
	fmt.Printf("First rune: %c %v bytes", c, size)

}

/*
** Playing with runes
** This is the maximum Strings in unsigned int 8 bit, 16 bit, 32 bit integers
 */

func lesson960() {
	var x32 uint32 = 4294967295
	var x16 uint16 = 65535
	var x8 uint8 = 255

	fmt.Printf("x32 %b, x16 %b, x8 %b", x32, x16, x8)

}

/*
** ROT13 (rotate 13) is a 20th century variant of Caesar cipher. It has one difference: it adds 13 instead of 3. With ROT13, ciphering and deciphering
** are the same convenient operation.
**
**
** Let’s suppose, while scanning the heavens for alien communications, the SETI Institute received a transmission with the following message_original:
**
** message_original := "uv vagreangvbany fcnpr fgngvba"
** We suspect this message_original is actually English text that was ciphered with ROT13. Call it a hunch. Before you can crack the code, there’s one more thing you need to know. This message_original is 30 characters long, which can be determined with the built-in len function:
**
** fmt.Println(len(message_original))       1
 */

func lesson950() {

	message_original := "uv vagreangvbany fcnpr fgngvba"

	for i := 0; i < len(message_original); i++ {
		c := message_original[i]
		if c >= 'a' && c <= 'z' {
			c = c + 13
			if c > 'z' {
				c = c - 26
			}
		}
		fmt.Printf("%c", c)
	}
}

/**
 ** Experiment: piggy.go
 ** Write a new piggy bank program that uses integers to track the number of cents rather than dollars. Randomly place nickels (5¢), dimes (10¢), and
 ** quarters (25¢) into an empty piggy bank until it contains at least $20.
 **
 ** Display the running balance of the piggy bank after each deposit in dollars (for example, $1.05).
 **
 ** Tip: If you need to find the remainder of dividing two numbers, use modulus (%).
 **/

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

/**
**
** Lesson 5. Capstone: Ticket to Mars
** Welcome to the first challenge. It’s time to take everything covered in unit 1 and write a program on your own. Your challenge is
** to write a ticket generator in the Go Playground that makes use of variables, constants, switch, if, and for. It should also draw on
** the fmt and math/rand packages to display and align text and to generate random numbers.
**
** When planning a trip to Mars, it would be handy to have ticket pricing from multiple spacelines in one place. Websites exist that
** aggregate ticket prices for airlines, but so far nothing exists for spacelines. That’s not a problem for you, though. You can use
** Go to teach your computer to solve problems like this.
**
**
**
** Start by building a prototype that generates 10 random tickets and displays them in a tabular format with a nice header, as follows:
**
** Spaceline        Days Trip type  Price
** ======================================
** Virgin Galactic    23 Round-trip $  96
** Virgin Galactic    39 One-way    $  37
** SpaceX             31 One-way    $  41
** Space Adventures   22 Round-trip $ 100
** Space Adventures   22 One-way    $  50
** Virgin Galactic    30 Round-trip $  84
** Virgin Galactic    24 Round-trip $  94
** Space Adventures   27 One-way    $  44
** Space Adventures   28 Round-trip $  86
** SpaceX             41 Round-trip $  72
** The table should have four columns:
**
** The spaceline company providing the service
** The duration in days for the trip to Mars (one-way)
** Whether the price covers a return trip
** The price in millions of dollars
** For each ticket, randomly select one of the following spacelines: Space Adventures, SpaceX, or Virgin Galactic.
**
** Use October 13, 2020 as the departure date for all tickets. Mars will be 62,100,000 km away from Earth at the time.
**
** Randomly choose the speed the ship will travel, from 16 to 30 km/s. This will determine the duration for the trip to Mars and also
** the ticket price. Make faster ships more expensive, ranging in price from $36 million to $50 million. Double the price for round
** trips.
**
** When you’re done, post your solution to the Get Programming with Go forums at forums.manning.com/forums/get-programming-with-go. If
** you get stuck, feel free to ask questions on the forums, or take a peek at the appendix for our solution.
**
**/

func lesson500() {

	const travelDistance = 62100000 // KM

	fmt.Printf("Spaceline        Days Trip type    Price\n")
	fmt.Printf("========================================\n")

	for count := 0; count < 10; count++ {

		travelSpeed := rand.Intn(15) + 16
		travelDays := ((travelDistance / travelSpeed) / 360) / 24
		//fmt.Printf("%v km/s %v days\n", travelSpeed, travelDays)
		travelPrice := travelSpeed + 20

		switch rand.Intn(3) {
		case 0:
			fmt.Printf("%-18v", "Virgin Galactic")
		case 1:
			fmt.Printf("%-18v", "SpaceX")
		default:
			fmt.Printf("%-18v", "Space Adventures")
		}

		fmt.Printf("%-4v", travelDays)

		switch rand.Intn(2) {
		case 0:
			fmt.Printf("%-11v $  %2v\n", "Round-trip", 2*travelPrice)
		case 1:
			fmt.Printf("%-11v $  %2v\n", "one-way", travelPrice)
		}
	}
}

//Generate a random year instead of always using 2018.
//For February, assign daysInMonth to 29 for leap years and 28 for other years.
//Hint: you can put an if statement inside of a case block.
//Use a for loop to generate and display 10 random dates.
func lesson414() {

	era := "AD"
	daysInMonth := 31
	currentYear := 2020

	for i := 11; i > 0; i-- {
		day := rand.Intn(daysInMonth) + 1
		year := rand.Intn(currentYear) + 1
		month := rand.Intn(12) + 1

		switch month {
		case 2:

			if (year % 4) == 0 {
				daysInMonth = 29
			} else {
				daysInMonth = 28
			}

		case 4, 6, 9, 11:
			daysInMonth = 30
		}

		fmt.Println(era, year, month, day)
	}
}

// Write a guess-the-number program. Make the computer pick random numbers between 1–100 until it
// guesses your number, which you declare at the top of the program. Display each guess and whether it
// was too big or too small.
func lesson413() {
	var myNumber = 12
	var yourNumber = 0
	var youGotIt = false

	for !youGotIt {
		time.Sleep(time.Second)
		yourNumber = rand.Intn(100) + 1
		fmt.Printf("\nYou guessed %v", yourNumber)
		if yourNumber == myNumber {
			youGotIt = true
			fmt.Printf(" that was correct!")
		}

	}
}
