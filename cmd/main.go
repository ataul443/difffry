package main

import (
	"fmt"
	"os"

	"github.com/ataul443/diffry"
)

func main() {
	fmt.Printf("Input both strings seperated by space -\n")
	var a, b string

	_, err := fmt.Scanf("%s %s", &a, &b)
	if err != nil {
		fmt.Printf("invalid input: %s", err.Error())
		os.Exit(1)
	}

	// diffry.Diff(a, b)
	diffry.Diff(a, b)
}
