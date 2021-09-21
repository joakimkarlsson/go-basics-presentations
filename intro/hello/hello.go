package main

import "fmt"

type User struct { // Available outside of package
	FirstName string // Available outside of package
	LastName  string
	canGreet  bool // Unavailable outside of package
}

func main() {
	user := User{
		FirstName: "Ellen",
		LastName:  "Ripley",
		canGreet:  true,
	}

	fmt.Printf("Hello, %s\n", user)
}
