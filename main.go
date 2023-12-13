package main

import (
	"fmt"
	"os"
	"spirosbond/shred"
)
	func main() {

	args := os.Args[1:] // Args without program name
	n_of_args := len(args)
	if n_of_args != 1{
		fmt.Println("Error! Expecting 1 argument, got: ", n_of_args)
		return
	}
	path := args[0]
	fmt.Println("Preparing to Shred:", path)
	res := shred.Shred(path)
	
	if res != nil{
		panic(res)
	}
	fmt.Println("Shred result:", res)

}  
