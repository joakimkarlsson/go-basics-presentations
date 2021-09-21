package main

import "fmt"

type User struct {
	FirstName string
	LastName  string
	canGreet  bool
}

func (u User) String() string { // HL
	return fmt.Sprintf("%s %s", u.FirstName, u.LastName) // HL
} // HL

func main() {
	user := User{
		FirstName: "Ellen",
		LastName:  "Ripley",
		canGreet:  true,
	}

	fmt.Printf("Hello, %s\n", user)
}
