package main

import "fmt"

// Generic Exam struct without type constraints
type Exam[T any] struct {
	ExamCode T
	Subject  string
	Marks    int
}

func main() {

	// Exam with int ExamCode
	exam1 := Exam[int]{ExamCode: 101, Subject: "Math", Marks: 95}
	fmt.Printf("Exam1: %+v\n", exam1)

	// Exam with float64 ExamCode
	exam2 := Exam[float64]{ExamCode: 202.5, Subject: "Physics", Marks: 88}
	fmt.Printf("Exam2: %+v\n", exam2)

	// Exam with string ExamCode
	exam3 := Exam[string]{ExamCode: "CS-303", Subject: "Computer Science", Marks: 92}
	fmt.Printf("Exam3: %+v\n", exam3)

}
