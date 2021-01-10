/**
** To send ciphered messages, write a program that ciphers plain text using a keyword:
**
** plainText := "your message goes here"
** keyword := "GOLANG"
** Bonus: rather than write your plain text message in uppercase letters with no spaces, use the strings.Replace and strings.ToUpper functions to remove   ** spaces and uppercase the string before you cipher it.
**
** Once you’ve ciphered a plain text message, check your work by deciphering the ciphered text with the same keyword.
**/

package practicefundamentals

import (
	"fmt"
	"strings"
	"unicode/utf8"
)

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

	messageOriginal := "Hola Estación Especial Internacional"
	var messageNew string
	var messageNewWithRange string

	// fmt.Printf("Tip: ASCII String of %c is %[1]v\n", 'o')
	// fmt.Printf("Tip: ASCII String of %c is %[1]v\n", 'ó')
	// fmt.Printf("Tip: ASCII String of %c is %[1]v\n", 'a')
	// fmt.Printf("Tip: ASCII String of %c is %[1]v\n", 'A')
	// fmt.Printf("Tip: ASCII String of %c is %[1]v\n", 'z')
	// fmt.Printf("Tip: ASCII String of %c is %[1]v\n", 'Z')

	fmt.Printf("\nThe original Message was:       %v\n", messageOriginal)

	//fmt.Printf("\n\n (##) O NR\n")

	for _, c := range messageOriginal {
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
		messageNewWithRange += string(c)
	}

	//fmt.Printf("Processing the For Loop:\n\n")
	//fmt.Printf("(##) O N\n")

	for i := 0; i < len(messageOriginal); i++ {
		c := messageOriginal[i]
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
		messageNew += string(c)
		//mt.Printf("%c", c)
	}

	for i := 0; i < len(messageOriginal); i++ {
		//fmt.Printf("\n(%v) Looks like the charector %c in original message", i, messageOriginal[i])
		//fmt.Printf("\nIs being turned to %c in new message   ", message_new[i])
		//fmt.Printf("\nIs being turned to %c in Range message \n", message_new_with_range[i])
	}

	fmt.Printf("The new with ranges message is: %v\n", messageNewWithRange)
	fmt.Printf("The new message is:             %v\n", messageNew)

}

/*
** ROT13 (rotate 13) is a 20th century variant of Caesar cipher. It has one difference: it adds 13 instead of 3. With ROT13, ciphering and deciphering
** are the same convenient operation.
**
**
** Let’s suppose, while scanning the heavens for alien communications, the SETI Institute received a transmission with the following messageOriginal:
**
** messageOriginal := "uv vagreangvbany fcnpr fgngvba"
** We suspect this messageOriginal is actually English text that was ciphered with ROT13. Call it a hunch. Before you can crack the code, there’s one more thing you need to know. This messageOriginal is 30 characters long, which can be determined with the built-in len function:
**
** fmt.Println(len(messageOriginal))       1
 */

func lesson950() {

	messageOriginal := "uv vagreangvbany fcnpr fgngvba"

	for i := 0; i < len(messageOriginal); i++ {
		c := messageOriginal[i]
		if c >= 'a' && c <= 'z' {
			c = c + 13
			if c > 'z' {
				c = c - 26
			}
		}
		fmt.Printf("%c", c)
	}
}
