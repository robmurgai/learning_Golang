package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

func getInput() (magazine []string, note []string) {

	// Read the file line by line
	dir := "hackerrank"
	fileName := "checkMagazing_input16.txt"
	filePath := filepath.Join(dir, fileName)
	fmt.Printf("DEBUG: filePath: %v\n", filePath)

	file, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scannedLine := bufio.NewScanner(file)
	buf := make([]byte, 0, 64*1024)
	scannedLine.Buffer(buf, 1024*1024)

	//Scan the First line, it is m and n integers denoting the size of magazing and notes.
	scannedLine.Scan()
	//fmt.Println("Scanned Line: ", scannedLine.Text())
	lineSplits := strings.Split(scannedLine.Text(), " ")
	magSize, _ := strconv.Atoi(lineSplits[0])
	fmt.Printf("DEBUG: magSize: %v\n", magSize)

	noteSize, _ := strconv.Atoi(lineSplits[1])
	fmt.Printf("DEBUG: noteSize: %v\n", noteSize)

	//Scan the Second line, it is strings in the magazing
	scannedLine.Scan()
	//fmt.Println("Scanned Line: ", scannedLine.Text())
	magazine = strings.Split(scannedLine.Text(), " ")
	fmt.Printf("DEBUG: magazine size: %d\n", len(magazine))

	//Scan the Third line, it is strings in the note
	scannedLine.Scan()
	//fmt.Println("Scanned Line: ", scannedLine.Text())
	note = strings.Split(scannedLine.Text(), " ")
	fmt.Printf("DEBUG: note size: %d\n", len(note))

	err = scannedLine.Err()
	if err != nil {
		log.Fatal(err)
	}

	return magazine, note

}
