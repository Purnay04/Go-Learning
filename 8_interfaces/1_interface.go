package main

import "fmt"

type Printable interface {
	print()
}

type Printer struct {
	name        string
	printerType string
}

// receiver param
func (printer Printer) print() {
	fmt.Printf("Printer=%s of type=%s has implemented Printable.\n", printer.name, printer.printerType)
}

func main() {

	p := Printer{"P-098", "Laser"}

	p.print()

}
