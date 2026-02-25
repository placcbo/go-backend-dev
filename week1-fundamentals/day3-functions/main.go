package main

import (
	"errors"
	"fmt"
)

func main() {
	fmt.Println(celsiusToFahrenheit(10.00))
	fmt.Println(kgToLbs(10.00))
	fmt.Println(kmToMiles(-10))

}
func celsiusToFahrenheit(c float64) float64 {
	return float64((c * 9 / 5) + 32)

}

func kgToLbs(kg float64) (float64, error) {
	if kg < 0 {
		return 0.00, errors.New("kg cannot be zero")

	}
	return kg * 2.20462, nil

}

func kmToMiles(km float64) (float64, error) {
	if km < 0 {
		return 0.0, errors.New("km cannot be negative")
	}
	return km * 0.621371, nil
}
