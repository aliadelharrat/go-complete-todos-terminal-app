package output

import "fmt"

func PrintSuccess(msg string) {
	fmt.Printf("\033[32m%s\033[0m\n", msg)
}

func PrintError(msg string) {
	fmt.Printf("\033[31m%s\033[0m\n", msg)
}
