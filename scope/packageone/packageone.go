package packageone

import "fmt"

var PackageVar = "Package variable in packageone"

func PrintMe(myVar, blockVar string) {
	fmt.Println(myVar, blockVar, PackageVar)
}
