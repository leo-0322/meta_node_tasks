package main

type Person struct {
	Name string
	Age  int
}

type Employee struct {
	EmployeeID int
	Person
}

func (e Employee) PrintInfo() {
	println("Name:", e.Name)
	println("Age:", e.Age)
	println("EmployeeID:", e.EmployeeID)
}

func main() {
	e := Employee{EmployeeID: 12345, Person: Person{Name: "Alice", Age: 30}}
	e.PrintInfo()
}
