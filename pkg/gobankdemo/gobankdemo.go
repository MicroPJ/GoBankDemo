package gobankdemo

import (
	"fmt"
	"runtime"
)

func One(input []string) (result string) {

	fmt.Printf("*---[One] Running\n")
	switch runtime.GOOS {
	case "windows":
		fmt.Printf("*---[One] Windows identified\n")

	default: //Mac & Linux
		fmt.Printf("*---[One] Linux identified\n")
	}

	if len(input) == 0 {
		fmt.Printf("*---[One] No Parameter provided\n")
	} else {
		fmt.Printf("*---[One] Parameter provided: %v\n", input)
	}

	fmt.Printf("*---[One] Completed\n")
	return "Done"
}
