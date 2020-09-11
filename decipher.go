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

package main

import "fmt"

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
	messageOriginal := "L fdph, L vdz, L frqtxhuhg."
	var messageNew string

	fmt.Printf("\nThe original Message was:  %v\n", messageOriginal)
	//fmt.Printf("Processing the For Loop:\n\n")

	for i := 0; i < len(messageOriginal); i++ {
		c := messageOriginal[i]
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
		messageNew += string(c)
		//fmt.Printf("%c", c)
	}

	fmt.Printf("The new message is:        %v\n", messageNew)

}
