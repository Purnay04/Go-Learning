package main

import (
	"fmt"
	"math"
)

func failureHandler(msg string) {
	panic(msg)
}

func printBookPageDetails(pgNum int, offset int, books []string) {
	fmt.Printf("Books on page %d:", pgNum)
	fmt.Println()
	for idx, val := range books {
		fmt.Printf("%d: %s", offset+idx+1, val)
		fmt.Println()
	}
}

func getBookPageData(books []string, pgSize, pgNum int, failureHandler func(msg string)) (int, int, []string) {
	totPg := math.Ceil(float64(len(books)) / float64(pgSize))
	if pgNum > int(totPg) {
		failureHandler("Page is not present")
	}

	if pgNum == int(totPg) {
		offset := (pgNum - 1) * pgSize
		return pgNum, offset, books[offset:]

	} else {
		offset := (pgNum - 1) * pgSize
		return pgNum, offset, books[offset : offset+pgSize]
	}
}

func main() {
	var books [10]string = [10]string{
		"Book One",
		"Book Two",
		"Book Three",
		"Book Four",
		"Book Five",
		"Book Six",
		"Book Seven",
		"Book Eight",
		"Book Nine",
		"Book Ten",
	}

	fmt.Println("Available Books: ", books)

	var pgSize, pgNum int

	fmt.Print("Enter books per page:")
	fmt.Scan(&pgSize)

	fmt.Print("Enter page number:")
	fmt.Scan(&pgNum)

	printBookPageDetails(getBookPageData(books[:], pgSize, pgNum, failureHandler))
}
