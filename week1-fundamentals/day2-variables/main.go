package main

import "fmt"

func main() {
	 var hourlyRate float64 = 1.40

    // Way 2: var with inference
    var currentJob = "CloudFactory"

    // Way 3: short declaration (most common)
    targetSalary := 45000.0

    // Calculations
    weeklyEarnings  := hourlyRate * float64(hoursPerWeek)
    monthlyEarnings := weeklyEarnings * 4
    annualEarnings  := weeklyEarnings * 52

    // Years to earn target salary at current rate
    yearsNeeded := targetSalary / annualEarnings

    // Print it all out
    fmt.Println("=== Earnings Report ===")
    fmt.Printf("Job:              %s\n", currentJob)
    fmt.Printf("Hourly rate:      $%.2f\n", hourlyRate)
    fmt.Printf("Weekly earnings:  $%.2f\n", weeklyEarnings)
    fmt.Printf("Monthly earnings: $%.2f\n", monthlyEarnings)
    fmt.Printf("Annual earnings:  $%.2f\n", annualEarnings)
    fmt.Printf("Target salary:    $%.2f\n", targetSalary)
    fmt.Printf("Years at CF to reach target: %.1f years\n", yearsNeeded)
    fmt.Println("Solution: Get out of CF in 20 weeks.")

}
