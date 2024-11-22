package main

import "fmt"

// x := 10 not allowed outside the method or global place
// var limit int = 300     //allowed
// var name = "Amar kumar" //allowed
// limit = 400             //not allowed
const LoginToken string = "obhdmhghhsdvhngdh" //public var

func main() {
	fmt.Println("hello world!")
	fmt.Println("Variables")
	fmt.Println("Data Types")

	// var name type
	var username string = "Amar@anand"
	fmt.Println("Username", username)
	fmt.Printf("Type of that var: %T \n", username)

	var isLoggedIn bool = true
	fmt.Println("Is logged in", isLoggedIn)
	fmt.Println("Login token =", LoginToken)
	fmt.Printf("Type of LoginToken var: %T \n", LoginToken)
	fmt.Printf("Type of that var: %T \n", isLoggedIn)

	// Multiple variables
	var a int = 10
	var b int = 20
	var c int = a + b
	fmt.Println("Sum", c)
	fmt.Printf("Type of that var: %T \n", c)

	// Short hand declaration using := opreator
	x := 10
	y := 20
	z := x + y
	fmt.Println("Sum", z)
	fmt.Printf("Type of that var: %T \n", z)

	var p uint8 = 10
	var q uint8 = 20
	var r uint8 = p + q
	fmt.Println("Sum", r)
	fmt.Printf("Type of that var: %T \n", r)

	var num int
	fmt.Println("Nunber default value", num)

	// Type conversion
	var s int = 10
	var t float64 = 20.5
	var u float64 = float64(s) + t
	fmt.Println("Sum", u)

	// Taking input from user
	var name string
	fmt.Scanln(&name)
	fmt.Println("Your name is", name)

	// Constants
	const pi float64 = 3.14159265359
	fmt.Println("Value of pi", pi)
	fmt.Printf("Type of that var: %T \n", pi)

}
