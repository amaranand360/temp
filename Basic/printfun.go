package main

import "fmt"

func printfun() {
	var1 := "Hello"
	var2 := 42
	var3 := 3.14
	fmt.Printf("%v %v %v\n", var1, var2, var3) // // Prints: Hello 42 3.14
	fmt.Printf("%v %v %v\n", var1, var2, var3) // // Prints: Hello 42 3.14

	fmt.Printf("%T %T %T\n", var1, var2, var3) // Type of variables

	fmt.Printf("var1: %v (%T), var2: %v (%T), var3: %v (%T)\n", var1, var1, var2, var2, var3, var3)
	// Output: var1: Hello (string), var2: 42 (int), var3: 3.14 (float64)

	fmt.Printf("First variable: \"%v\", Second variable: \"%v\", Third variable: \"%v\"\n", var1, var2, var3)
	// Output: First variable: "Hello", Second variable: "42", Third variable: "3.14"

	fmt.Printf("String: %s, Integer: %d, Float: %.2f\n", var1, var2, var3)
	// Output: String: Hello, Integer: 42, Float: 3.14

	formattedString := fmt.Sprintf("String: %s, Integer: %d, Float: %.2f", var1, var2, var3)
	fmt.Println(formattedString)
	// Output: String: Hello, Integer: 42, Float: 3.14

}
