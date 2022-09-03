package utils

import "fmt"

var (
	CncAddr    = "192.168.1.15"
	CncPort    = 444
	CncApiPort = 3333

	SingleInstancePort = 1337
	BinVersion         = "0.0.3"
	DebugEnabled       = true

	Edpoint = map[string]string{
		"update": fmt.Sprintf("http://%s:%d/update", CncAddr, CncApiPort),
	}
)

var (
	InstanceRunning = true
)
