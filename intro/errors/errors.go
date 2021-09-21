package main

import (
	"errors"
	"fmt"
)

type User struct {
	FirstName string
	LastName  string
	canGreet  bool
}

func (u User) String() string { // HL
	return fmt.Sprintf("%s %s", u.FirstName, u.LastName) // HL
} // HL

// START OMIT
func makeGreeting(u User) (string, error) {
	if !u.canGreet {
		return "", errors.New("cannot greet")
	}

	return fmt.Sprintf("Hello, %s %s", u.FirstName, u.LastName), nil
}

// END OMIT

// START2 OMIT
func main() {
	// OMIT
	user := User{ // OMIT
		FirstName: "Ellen",  // OMIT
		LastName:  "Ripley", // OMIT
		canGreet:  true,     // OMIT
	} // OMIT

	greeting, err := makeGreeting(user)
	if err != nil {
		fmt.Print("cannot greet: ", err)
		return
	}

	fmt.Println(greeting)
}

// END2 OMIT
