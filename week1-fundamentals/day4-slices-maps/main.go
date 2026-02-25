package main

import (
	"fmt"
)

func main() {
words := [] string {"apple", "banana", "cherry", "apple", "date", "banana", "fig", "grape", "cherry", "apple"}
freq := make(map[string] int)



// count each word 

for _, word := range words {
	freq[word]++
}

//  print frequencies

for word, count := range freq{
	fmt.Printf("%s: %d\n", word, count)
}
}
