package main

import "fmt"

func main() {
	var name string

	name = "Iqbal"
	fmt.Println(name)

	name = "Nugraha"
	fmt.Println(name)

	var friendName = "Budi"
	fmt.Println(friendName)

	var age = 25
	fmt.Println(age)

	// Multiple Var
	var (
		firstName = "Muhammad"
		lastName  = "Iqbal"
	)

	fmt.Println(firstName + " " + lastName)
}
