package main

import "fmt"

func main() {
	even := 6

	//if
	if even == 6 {
		fmt.Printf("Number is %d\n", even)
	}

	//else if , else blocks
	age := 34
	if age < 12 {
		fmt.Println("Age is less than 12 years")
	} else if age >= 12 && age <= 24 {
		fmt.Println("Age is greater than or equal to 12 and less than 24")
	} else {
		fmt.Println("Age is greater than or equal to 24")
	}

}
