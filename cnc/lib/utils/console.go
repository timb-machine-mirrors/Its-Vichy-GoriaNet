package utils

import "fmt"

func Log(Content string) {
	fmt.Println(fmt.Sprintf("%s.", Content))
}

func Debug(Content string) {
	fmt.Println(fmt.Sprintf("%s.", Content))
}
