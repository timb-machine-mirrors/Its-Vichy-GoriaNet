package security

import (
	"syscall"
)

func SafeExit() {
	defer syscall.Exit(0)

	ExecuteGroup([]string{
		"rm -rf /var/tmp/*",
		"iptables -F",
	})
}