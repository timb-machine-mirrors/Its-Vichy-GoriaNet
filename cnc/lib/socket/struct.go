package socket

import "net"

var (
	BotList    []*Bot
	MasterList []*Master
)

type BotSocket struct {
	Connected bool
	Socket    net.Conn
}

type Bot struct {
	Cpu     int
	Mem     int
	Disk    string
	Arch    string
	Version string
	Network BotSocket
}

type MasterSocket struct {
	Connected bool
	Socket    net.Conn
}

type Master struct {
	Network MasterSocket
}

func RemoveBot(B *Bot) {
	for i, v := range BotList {
		if v == B {
			BotList = append(BotList[:i], BotList[i+1:]...)
		}
	}
}
