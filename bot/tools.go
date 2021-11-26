package main

import "fmt"

func debug(content string) {
	if debug_mode {
		fmt.Printf("[DEBUG] %s\n", content)
	}
}
