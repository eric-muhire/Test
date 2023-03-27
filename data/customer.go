package data

var antal int

func getNrOfCustomers() int {
	return antal
}

type Customer struct {
	Name string
	Age  int
}

func GetAllCustomers() []Customer {
	var allCustomerSlice = []Customer{}
	allCustomerSlice = append(allCustomerSlice, Customer{Name: "Anna", Age: 50})
	allCustomerSlice = append(allCustomerSlice, Customer{Name: "marie", Age: 70})
	allCustomerSlice = append(allCustomerSlice, Customer{Name: "per", Age: 60})

	return allCustomerSlice
}
