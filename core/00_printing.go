package main

import "fmt"

func main() {
	fmt.Print("Hello World", 12, "\n")
	fmt.Print("Ironman is dead", "So sad of you", "\n")

	fmt.Println("Go is from Google")
	fmt.Println("Go is now opensourced")

	fmt.Printf("Hi! My name is %s. I have lived in %s for %d years. They say the weather is amazing, which is %t.\n", "Ironman", "LA", 4, true)

	// string formatting
	fmt.Printf("My name is : %s, and it is %t that i have %d brother\n", "Balley", true, 1)

}
