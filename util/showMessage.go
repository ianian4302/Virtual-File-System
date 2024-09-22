package util

import "fmt"

const (
	Reset = "\033[0m"
	Red   = "\033[31;1m"
	Green = "\033[32;1m"
)

func CallError(str string) {
	fmt.Println(Red, "Error: ", str, Reset)
}

func CallSuccess(str string) {
	fmt.Println(Green, str, Reset)
}
