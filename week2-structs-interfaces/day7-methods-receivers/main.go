package main

import (
	"encoding/json"
	"fmt"
)

type Company struct {
	Name     string
	Location Address
}

type Address struct {
	City    string
	Country string
}
type JobApplication struct {
	ID        int     `json:"id"`
	Role      string  `json:"role"`
	Company   Company `json:"company"`
	SalaryUSD int     `json:"salaryusd"`
	Status    string  `json:"status"`
	AppliedAt string  `json:"appliedat,omitempty"`
	IsRemote  bool    `json:"isremote"`
}

func main() {
	applications := []JobApplication{
		{
			ID:   1,
			Role: "backend dev",
			Company: Company{
				Name: "Cloudfactory",
				Location: Address{
					Country: "kenya",
					City:    "Nairobi",
				},
			},
			SalaryUSD: 300,
			Status:    "pending",
			AppliedAt: "",
			IsRemote:  true,
		},
		{
			ID:   2,
			Role: "QA engineer",
			Company: Company{
				Name: "nuvita",
				Location: Address{
					Country: "Uganda",
					City:    "kamplala",
				},
			},
			SalaryUSD: 500,
			Status:    "pending",
			AppliedAt: "11/11/2023",
			IsRemote:  false,
		},
		{
			ID:   3,
			Role: "backend engineer",
			Company: Company{
				Name: "youtube",
				Location: Address{
					Country: "rwanda",
					City:    "Kigali",
				},
			},
			SalaryUSD: 300,
			Status:    "rejected",
			AppliedAt: "2025-11-09",
			IsRemote:  false,
		},
	}
	data, err := json.MarshalIndent(applications, "", "")
	if err != nil {
		fmt.Println("Error: ", err)
	}
	fmt.Println(string(data))

	filterByPending := filterByStatus(applications, "rejected")
	fmt.Println("filter by pending: ", filterByPending)
}

func filterByStatus(apps []JobApplication, status string) []JobApplication {
	results := make([]JobApplication, 0)
	for _, app := range apps {
		if app.Status == status {
			results = append(results, app)
		}

	}
	return results
}
