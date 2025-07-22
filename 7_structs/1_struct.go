package main

type Employee struct {
	id     string
	name   string
	salary float32
}

func main() {

	// simple employee var
	// fmt.Println(Employee{"QWA-2435", "John Doe", 12345.6})

	// fields can be accessed and mutated by a dot operator
	// e1 := Employee{"QWA-2435", "John Doe", 12345.6}
	// fmt.Println(e1.name)

	// pointer to struct var
	// ptr *Employee = &e1
	// p := &e1
	// e1.salary = 122345.6
	// (*p).salary = 122345.6
	// (*p).salary is tedious in usage, we can ignore the usage of *
	// p->salary = 122345.6 in c prog
	// p.salary = 122345.6
	// fmt.Println(e1)

	// Struct Literals

	// provide all values
	// e1 := Employee{"QWA-2435", "John Doe", 12345.6}

	// ignore name and salary
	// e1 := Employee{id: "QWA-2435", name: "john doe"}
	// fmt.Println(e1)

	// e1 := Employee{"QWA-2435", "John Doe", 12345.6}
	// e2 := Employee{"QWA-5431", "Mary K.", 17645.6}

	// emps := []Employee{e1, e2, Employee{"QOA-2435", "Roy J.", 42345.6}}
	// fmt.Println(emps)
}
