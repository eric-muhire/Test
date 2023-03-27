package main

import (
	"eric/education/data"
	"fmt"
)

func main() {
	fmt.Println("Hello")

	antal, average := data.GetStats()
	fmt.Printf("Antal kunder: %d %f\n", antal, average)

	//var nrOfCustomers int
	nrOfCustomers := data.GetNrOfCustomers()

	for _, customer := range data.GetAllCustomers() {
		fmt.Printf("kund nr:%d år\n, index+1")
		fmt.Printf("%s är %d år\n", customer.Name, customer.Age)
	}

	//antal, average :=data.getStats()
	fmt.Printf("Antal kunder: %d %f\n", nrOfCustomers)

}
