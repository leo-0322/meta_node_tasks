package main

import (
	"fmt"

	"github.com/jmoiron/sqlx"
)

type Employee struct {
	ID         int     `db:"id"`
	Name       string  `db:"name"`
	Department string  `db:"department"`
	Salary     float64 `db:"salary"`
}

var db *sqlx.DB

func findDepartmentInfo(department string) ([]Employee, error) {
	var employees []Employee

	query := "select id, name, department, salary from employees where department = ?"
	err := db.Select(&employees, query, department)
	if err != nil {
		return nil, err
	}
	return employees, nil
}

func findMaxSalary() (Employee, error) {
	var employee Employee
	query := "select id, name, department, salary from employees order by salary desc limit 1"
	err := db.Get(&employee, query)
	if err != nil {
		return Employee{}, err
	}
	return employee, nil
}

func main() {
	emps, err := findDepartmentInfo("技术部")
	fmt.Println("技术部员工信息：", emps, err)

	maxSalaryEmp, err := findMaxSalary()
	fmt.Println("Max Salary:", maxSalaryEmp, err)
}
