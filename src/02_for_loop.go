package main

import "fmt"

func main(){

	// c style for loop
	for count := 1 ; count < 10 ; count++ {
		fmt.Println("Count is : ", count)
	}


	// for loop with only condition
	var sum int = 0
	for sum < 22{
		sum += 3
		fmt.Println("Current sum is:", sum)

	}


	// for range 
	primes := []int{2, 3 , 5, 7, 11}
	for index, value := range primes{
		fmt.Println("Index is: ", index, "and value is:", value)


	}



	// use _ to ignore index
	for _, value := range primes{
		fmt.Println(value)
	}
}