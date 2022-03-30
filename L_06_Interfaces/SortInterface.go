package main

import (
	"fmt"
	"sort"
)

type Employee struct {
	Name string
	Age  int
}

type EmployeeCompare struct {
	Comparator func(first, second *Employee) bool
	Employee   []*Employee
}

func (emp *EmployeeCompare) Len() int {
	return len(emp.Employee)
}

func (emp *EmployeeCompare) Less(i, j int) bool {
	return emp.Comparator(emp.Employee[i], emp.Employee[j])
}

func (emp *EmployeeCompare) Swap(i, j int) {
	emp.Employee[i], emp.Employee[j] = emp.Employee[j], emp.Employee[i]
}

func byAge(first *Employee, second *Employee) bool {
	return first.Age < second.Age
}

func byName(first *Employee, second *Employee) bool {
	return first.Name < second.Name
}

func main() {

	employees := []*Employee{
		{"Suraj", 44},
		{"Chand", 34},
		{"Guru", 23},
		{"Ravi", 44},
		{"Soma", 31},
	}

	fmt.Println("Sorting by age")
	sortEmployee(employees, byAge)

	fmt.Println("Sorting by name")
	sortEmployee(employees, byName)

	fmt.Println("Sorting by age then name")
	sortEmployee(employees, func(first *Employee, second *Employee) bool {
		if first.Age != second.Age {
			return first.Age < second.Age
		}
		return first.Name < second.Name
	})

}

func sortEmployee(employees []*Employee, comparator func(first, second *Employee) bool) {

	comparable := EmployeeCompare{comparator, employees}

	//NOTE: Passing address of EmployeeCompare because
	sort.Sort(&comparable) //  (* EmployeeCompare) satisfies the interface and not EmployeeCompare

	for _, emp := range employees {
		fmt.Println(emp.Name, emp.Age)
	}
}
