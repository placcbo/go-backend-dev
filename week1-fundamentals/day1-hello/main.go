package main

import "fmt"

func main() {
	name := "Alice"
	age := 30
	salary := 1.40

	fmt.Println("name", name)
	fmt.Println("age", age)
	fmt.Println("salary", salary)

	fmt.Printf("name: %s\n", name)
	fmt.Printf("age: %d\n", age)
	fmt.Printf("salary: %.2f", salary)
	fmt.Printf("Type %T\n", salary)

	message := fmt.Sprintf("Hello, %s! You are %d.", name, age)
	fmt.Println(message)

}
