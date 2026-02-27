package main

import "fmt"

//word frequency counter

func main() {

	cities := []string{
		"Buenos Aires",
		"Rio de Janeiro",
		"Santiago",
		"Santiago",
		"Santiago",
		"Santiago",
		"Santiago",
		"Santiago",
		"Santiago",
		"Santiago",
		"Lima",
		"Santiago",
		"Santiago",
		"Santiago",
		"Santiago",
		"Bogotá",
		"Buenos Aires",
		"Quito",
		"Lima",
		"São Paulo",
		"São Paulo",
		"São Paulo",
		"São Paulo",
		"São Paulo",
		"Caracas",
		"Rio de Janeiro",
		"Rio de Janeiro",
		"Rio de Janeiro",
	}
	mappedCity:= make(map[string]int)
	for _, city := range cities {
	mappedCity[city]++
	}
	result := topWords(mappedCity, 1)
	fmt.Println(result)
	

}

func topWords(freq map[string]int, minCount int) []string {
	newCitySlice := []string{}
	for word, count := range freq {
		if count > minCount {
			newCitySlice = append(newCitySlice, word)
		}
	}
	return newCitySlice
}
