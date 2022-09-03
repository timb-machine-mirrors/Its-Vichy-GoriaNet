package utils

import (
	"fmt"
)

func HandleError(err error) {
	if err != nil {
		fmt.Println(err)
		return
	}
}

func GetTitle(content string) string {
	return fmt.Sprintf("\033]0; %s\007", content)
}

func GetRGB(R int, G int, B int) string {
	return fmt.Sprintf("\033[38;2;%d;%d;%dm", R, G, B)
}

func GetClear() string {
	return "\033[2J\033[1H"
}
