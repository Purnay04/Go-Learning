package main

import "fmt"

type Certificate struct {
	id     string
	name   string
	issuer string
}

type Employee struct {
	id           string
	name         string
	salary       float32
	certificates []Certificate
}

func main() {

	employees := []Employee{Employee{"QOA-2435", "Roy J.", 42345.6,
		[]Certificate{{"1", "AZ-204", "MS"}, {"2", "AZ-304", "MS"}, {"3", "AI-104", "MS"}, {"4", "GC-12", "Google"},
			{"5", "AWS-122", "Amazon"}}}}

	fmt.Println(employees)

}
