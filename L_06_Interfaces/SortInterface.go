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
	Comparator func(comparator *EmployeeCompare, i, j int) bool
	Employee   []*Employee
}

func (emp *EmployeeCompare) Len() int {
	return len(emp.Employee)
}

func (emp *EmployeeCompare) Less(i, j int) bool {
	return emp.Comparator(emp, i, j)
}

func (emp *EmployeeCompare) Swap(i, j int) {
	emp.Employee[i], emp.Employee[j] = emp.Employee[j], emp.Employee[i]
}

func (emp *EmployeeCompare) byAge(i, j int) bool {
	return emp.Employee[i].Age < emp.Employee[j].Age
}

func (emp *EmployeeCompare) byName(i, j int) bool {
	return emp.Employee[i].Name < emp.Employee[j].Name
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
	sortEmployee(employees, (*EmployeeCompare).byAge)

	fmt.Println("Sorting by name")
	sortEmployee(employees, (*EmployeeCompare).byName)

	fmt.Println("Sorting by age then name")
	sortEmployee(employees, func(emp *EmployeeCompare, i, j int) bool {
		if emp.Employee[i].Age != emp.Employee[j].Age {
			return emp.Employee[i].Age < emp.Employee[j].Age
		}

		return emp.Employee[i].Name < emp.Employee[j].Name
	})

}

func sortEmployee(employees []*Employee, comparator func(empCompare *EmployeeCompare, i, j int) bool) {

	comparable := EmployeeCompare{comparator, employees}

	//NOTE: Passing address of EmployeeCompare because
	sort.Sort(&comparable) //  (* EmployeeCompare) satisfies the interface and not EmployeeCompare

	for _, emp := range employees {
		fmt.Println(emp.Name, emp.Age)
	}
}
