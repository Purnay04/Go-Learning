package main

import (
	"errors"
	"fmt"
)

// Calculator interface defines basic arithmetic operations
type Calculator interface {
	Add(a, b float64) float64
	Subtract(a, b float64) float64
	Multiply(a, b float64) float64
	Divide(a, b float64) (float64, error)
}

// SimpleCalculator implements Calculator interface
type SimpleCalculator struct{}

func (c SimpleCalculator) Add(a, b float64) float64 {
	return a + b
}

func (c SimpleCalculator) Subtract(a, b float64) float64 {
	return a - b
}

func (c SimpleCalculator) Multiply(a, b float64) float64 {
	return a * b
}

func (c SimpleCalculator) Divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, errors.New("cannot divide by zero")
	}
	return a / b, nil
}

// CalculatorApp holds the Calculator dependency
type CalculatorApp struct {
	calc Calculator
}

// Constructor injection
func NewCalculatorApp(c Calculator) *CalculatorApp {
	return &CalculatorApp{calc: c}
}

func main() {
	calcService := SimpleCalculator{}
	app := NewCalculatorApp(calcService) // Inject calculator service

	a, b := 10.0, 5.0

	fmt.Printf("Add: %.2f + %.2f = %.2f\n", a, b, app.calc.Add(a, b))
	fmt.Printf("Subtract: %.2f - %.2f = %.2f\n", a, b, app.calc.Subtract(a, b))
	fmt.Printf("Multiply: %.2f * %.2f = %.2f\n", a, b, app.calc.Multiply(a, b))

	divResult, err := app.calc.Divide(a, b)
	if err != nil {
		fmt.Println("Divide error:", err)
	} else {
		fmt.Printf("Divide: %.2f / %.2f = %.2f\n", a, b, divResult)
	}
}
