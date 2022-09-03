package attack

type Attack struct {
	DestAddr string
	DestPort int
	Threads  int
	Running  *chan bool
	Payload  []byte
	Power    int
	Time     int
	Name     string
}

type Method struct {
	Name        string
	Payload     string
	PayloadSize int
}
