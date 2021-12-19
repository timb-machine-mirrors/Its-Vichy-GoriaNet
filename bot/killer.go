package main

import (
	"fmt"
	"net"
	"os"
	"os/exec"
	"time"
)

func bind_port(port int) bool {
	_, err := net.Listen("tcp", fmt.Sprintf("127.0.0.1:%d", port))

	if err != nil {
		return false
	} else {
		return true
	}
}

func bind_instance_port() {
	if !bind_port(single_instance_port) {
		debug("Another instance was detected, exit..")
		os.Exit(1)
	} else {
		debug(fmt.Sprintf("Single instance port %d was successfully bind", single_instance_port))
	}
}

// Not optimized for IoT btw work on vps: https://youtu.be/5zkFm_8-sPQ
func kill_all_by_port() {
	for port := 0; port != 65536; port++ {
		if port != single_instance_port {
			if !bind_port(port) {
				exec.Command("bash", "-c", fmt.Sprintf("lsof -i tcp:%d | grep LISTEN | awk '{print $2}' | xargs kill -9", port)).Run()

				if bind_port(port) {
					debug(fmt.Sprintf("Bind a killed port --> %d", port))
				}
			}
		}

		time.Sleep(100 * time.Millisecond) // Prevent "to many open file" (optimisation soon..)
	}
}
