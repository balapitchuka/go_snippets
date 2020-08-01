package main

import "fmt"

type User struct {
	ID          int
	name, email string
}

func main() {
	var u User
	fmt.Println(u)

	u.ID = 34
	u.name = "Ironman"
	u.email = "abc@gmail.com"
	fmt.Println(u)

	userKay := User{ID: 1, name: "Kay Ole", email: "kayl.il@gmail.com"}
	fmt.Println(userKay)

	userDescription := describeUser(userKay)
	print(userDescription)
}

func describeUser(u User) string {
	description := fmt.Sprintf("Name: %s, Id : %d, Email : %s", u.name, u.ID, u.email)
	return description
}
