package main

import (
	"encoding/json"
	"fmt"
	"strings"
)

//EMPLOYEE DIRECTORY

type Address struct {
	Street  string `json:"street"`
	City    string `json:"city"`
	Country string `json:"country"`
}
type Employee struct {
	ID         int     `json:"id"`
	Name       string  `json:"name"`
	Department string  `json:"department,omitempty"`
	Salary     float64 `json:"salary"`
	Address    Address
	IsManager  bool `json:"ismanager"`
}



func main() {
	employees := []Employee{
		{
			ID:         1,
			Name:       "kevin",
			Department: "extraction",
			Salary:     30000,
			Address: Address{
				Street:  "mowlem",
				City:    "Nairobi",
				Country: "Kenya",
			},
			IsManager: false,
		},
		{
			ID:         1,
			Name:       "Prajwal",
			Department: "lead",
			Salary:     70000,
			Address: Address{
				Street:  "Kathmandu",
				City:    "Kathmandu-Boga",
				Country: "Nepal",
			},
			IsManager: true,
		},
		{
			ID:         1,
			Name:       "Abigael",
			Department: "extraction",
			Salary:     32000,
			Address: Address{
				Street:  "Kanduyi",
				City:    "Bungoma",
				Country: "Kenya",
			},
			IsManager: false,
		},
	}
	 data, err := json.MarshalIndent(employees, "","")

	if err != nil{
		fmt.Println("Error: ", err)
	}
	fmt.Println(string(data))
		
	for _, e := range employees {
		dept := e.Department
		if dept == "" {
			dept = "No Department"
		}
		manager := ""
		if e.IsManager {
			manager = "Manager"
		}
		fmt.Printf("  %-16s %-16s KES %8.0f  %s%s\n",
			e.Name, dept, e.Salary, e.Address.City, manager)
	}

}
