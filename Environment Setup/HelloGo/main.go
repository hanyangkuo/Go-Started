package main

import (
	"fmt"
	"os"
)
func main(){
	if len(os.Args) < 2 {
		fmt.Printf("Usage: %s [Name]\n", os.Args[0])
		os.Exit(2)
	}
	fmt.Printf("Hello %s!\n", os.Args[1])
}