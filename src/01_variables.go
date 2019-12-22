package main

import "fmt"

func main() {
	fmt.Println("hello world")


	//  string
	var name string = "Ironman"
	fmt.Println("My name is :", name)
	



    // integer
	var height, weight int = 5, 62
	fmt.Println("My height is ", height, " feet.My weight is : ", weight, "kg")


	// boolean
	isDone := true
	var willRain = false
	var isWild bool = true
	fmt.Println("Is workdone? : ", isDone)
	fmt.Println("Will it rain tommorrow :", willRain)
	fmt.Println("Is cat wild :", isWild)

	//find types
	fmt.Printf("%T %T %T %T\n", name, height, isDone, willRain)

}
