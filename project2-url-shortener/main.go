package main

import (
	"fmt"
	"project2-url-shortener/database"
)

func main() {
	err := database.Connect()
	if err != nil {
		fmt.Println("error:", err)
		return
	}

	_, err = database.SaveURL("click77", "https://example.com")
	if err != nil {
		fmt.Println("save error:", err)
	}

	err = database.IncrementClicks("click77")
	if err != nil {
		fmt.Println("increment error:", err)
		return
	}

	err = database.IncrementClicks("click77")
	if err != nil {
		fmt.Println("increment error:", err)
		return
	}

	url, err := database.GetURL("click77")
	if err != nil {
		fmt.Println("get error:", err)
		return
	}

	fmt.Println("ShortCode:", url.ShortCode)
	fmt.Println("OriginalURL:", url.OriginalURL)
	fmt.Println("Clicks:", url.Clicks)
}

