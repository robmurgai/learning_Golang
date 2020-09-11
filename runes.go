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

package main

import (
	"fmt"
	"unicode/utf8"
)

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
