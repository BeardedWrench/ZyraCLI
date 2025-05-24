package main

import (
	"ZyraCLI/cmd"
	"fmt"
)
var version = "dev"

func main() {
	fmt.Println("ZyraCLI Version:", version)
	cmd.Execute()
}
