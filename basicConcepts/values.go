package main

import "fmt"

//values are datatypes in Go
func values() {
	fmt.Println("A" + "B")
	fmt.Println(1 + 1)
	fmt.Print(1.3 + 1.4)
	fmt.Println(true || false)

}

/*error PS C:\Users\Shiva.Karinigi\New folder\march-1\golang> go run values.go
# command-line-arguments
runtime.main_main·f: function main is undeclared in the main package*/

//Solved by adding main function which is the entry point of the gop program
