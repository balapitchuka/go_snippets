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
	var willRain = false //type is 	implied
	var isWild bool = true
	fmt.Println("Is workdone? : ", isDone)
	fmt.Println("Will it rain tommorrow :", willRain)
	fmt.Println("Is cat wild :", isWild)

	//find types
	fmt.Printf("%T %T %T %T\n", name, height, isDone, willRain)

	var x = "What am I"
	fmt.Printf("The type of `%v` is %T", x, x) //%v is a generic verb for "value", %T is a verb for Type

}
