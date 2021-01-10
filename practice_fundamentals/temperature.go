/**
** Lesson 15. Capstone: Temperature tables
** Write a program that displays temperature conversion tables. The tables should use equals signs (=) and pipes (|) to draw lines, with a header
** section:
**
**    =======================
**    | °C       | °F       |
**    =======================
**    | -40.0    | -40.0    |
**    | ...      | ...      |
**    =======================
** The program should draw two tables. The first table has two columns, with °C in the first column and °F in the second column. Loop from –40° C
** through 100° C in steps of 5° using the temperature conversion methods from lesson 13 to fill in both columns.
**
** After completing one table, implement a second table with the columns reversed, converting from °F to °C.
**
** Drawing lines and padding values is code you can reuse for any data that needs to be displayed in a two-column table. Use functions to separate the
** table drawing code from the code that calculates the temperatures for each row.
**
** Implement a drawTable function that takes a first-class function as a parameter and calls it to get data for each row drawn. Passing a different
** function to drawTable should result in different data being displayed.
**/

package practicefundamentals

import "fmt"

func lession15() {

	// Build a Celsius to Farenheit Table
	//Build the top row of the table with borders.
	fmt.Printf("%v\n", charLine("=", 23))
	fmt.Printf("%v", tableHeader("°C", "°F"))

	//Build the next set of rows with numbers.
	fmt.Printf("%v\n", charLine("=", 23))
	for i := -40.0; i < 101; i += 5 {
		drawTable(i, c2f)
	}

	//Close the table with the bottom Row
	fmt.Printf("%v\n", charLine("=", 23))

	fmt.Printf("\n")

	//Build a Farenheit to Celsius Table
	//Build the top row of the table with borders.
	fmt.Printf("%v\n", charLine("=", 23))
	fmt.Printf("%v", tableHeader("°F", "°C"))

	//Build the next set of rows with numbers.
	fmt.Printf("%v\n", charLine("=", 23))

	for i := -40.0; i < 101; i += 5 {
		drawTable(i, f2c)
	}

	//Close the table with the bottom Row
	fmt.Printf("%v\n", charLine("=", 23))

}

func charLine(c string, n int) string {

	var charString string

	for i := 0; i < n; i++ {
		charString += c
	}

	return charString
}

func spacedValue(value string, width int) string {
	var spacedString string

	spacedString = fmt.Sprintf(" %-9v", value)

	return spacedString

}

func tableHeader(value1 string, value2 string) string {

	var spacedString string
	spacedString = charLine("|", 1) + fmt.Sprintf(" %-9v", value1) + charLine("|", 1) + fmt.Sprintf(" %-9v", value2) + charLine("|", 1) + "\n"
	return spacedString

}

func tableRow(value1 float64, value2 float64) string {

	var spacedString string
	spacedString = charLine("|", 1) + fmt.Sprintf(" %-9.0f", value1) + charLine("|", 1) + fmt.Sprintf(" %-9.0f", value2) + charLine("|", 1) + "\n"
	return spacedString

}

func drawTable(tempOriginal float64, tempConverted func(float64) float64) {
	fmt.Printf("%v", tableRow(tempOriginal, tempConverted(tempOriginal)))
}

func c2f(i float64) float64 {
	return float64(celsius(i).farenheight())
}

func f2c(i float64) float64 {
	return float64(farenheight(i).celsius())
}
