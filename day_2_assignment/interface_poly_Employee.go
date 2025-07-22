package main

import (
	"fmt"
)

type Employee interface {
	CalculateSalary() float64
}

type PermanantEmployee struct {
	id    int
	name  string
	basic float64
	hra   float64
	da    float64
	ta    float64
}

type ContractEmployee struct {
	id          int
	name        string
	hours       int
	perHourRate float64
}

func (pEmp PermanantEmployee) CalculateSalary() float64 {
	return pEmp.basic + pEmp.hra + pEmp.da + pEmp.ta
}

func (cEmp ContractEmployee) CalculateSalary() float64 {
	return float64(cEmp.hours) * cEmp.perHourRate
}

func printEmployeeSalary(e Employee) {
	var eName string
	if pEmp, isSame := e.(PermanantEmployee); isSame {
		eName = pEmp.name
	} else if cEmp, isSame := e.(ContractEmployee); isSame {
		eName = cEmp.name
	}
	fmt.Printf("Salray of Employee %s: %.2f ", eName, e.CalculateSalary())
}

func main() {
	var employees []Employee = []Employee{
		PermanantEmployee{
			id:   1,
			name: "Purnay",
			hra:  20000.00,
			da:   10000.00,
			ta:   5000.00,
		},
		ContractEmployee{
			id:          2,
			name:        "Vinu",
			hours:       8,
			perHourRate: 800,
		},
	}

	for _, val := range employees {
		printEmployeeSalary(val)
		fmt.Println()
	}
}
