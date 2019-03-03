package main

import (
	"git.parallelcoin.io/tri"
	"fmt"
)

func main() {
	e := exampleTri.Validate()
	if e != nil {
		fmt.Println(e)
	}
	tri.LoadAllDefaults(&exampleTri)
}
