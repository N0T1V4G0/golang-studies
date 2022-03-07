package main

import "myapp/packageone"

var myVar = "Package level variable"

func main() {
	var blockVar = "Block level variable"
	packageone.PrintMe(myVar, blockVar)
}
