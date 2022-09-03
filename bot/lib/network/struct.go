package network

import (
	"bot/lib/utils"
	"runtime"
)

type BotInfo struct {
	Cpu  int
	Mem  int
	Disk string
	Arch string
}

var (
	Profile = BotInfo{
		Cpu:  runtime.NumCPU(),
		Mem:  utils.GetMemory(),
		Disk: utils.GetDiskSpace(),
		Arch: utils.GetCPUArch(),
	}
)
