package main

import "fmt"

func main() {
	var age int = 30
	var height, weight int = 30, 40
	
	var name = "ironman"
	address := "plot 301, highway, mars"
	var negativeNumber int = -459
	var sum int = -negativeNumber
	var negativeFloating = -459.12
	isDone := true

	fmt.Println("hello world")
	fmt.Println("My name is :", name)
	fmt.Println("My age is ", age)
	fmt.Println("My height is ", height)
	fmt.Println("My weight is ", weight)
	fmt.Println("My address is :", address)
	fmt.Println(negativeNumber, sum, negativeFloating)
	fmt.Println("Is workdone? : ", isDone)

	primes := []int{2, 3, 5, 7, 11}
	fmt.Println("primes numbers are :", primes)

	// appending to slice
	primes = append(primes, 13)

	fmt.Println("primes numbers are :", primes)

	//map
	capitals := map[string]string{"India": "Delhi", "UK": "Donot know"}
	fmt.Println(capitals)
}
