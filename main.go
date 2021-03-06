package main

var Flags FlagsConfig = NewFlagsConfig()

func main() {

	// Main channel for logs
	logs := make(chan LogMessage)

	config, err := NewConfig()
	if err != nil {
		Print(err)
		return
	}

	// Exit if new config was generated
	if config.new {
		Print("Empty configuration file was generated (config.json).")
		Print("Edit the file and re-run this command.")
		return
	}

	Print("Using %s configuration file.", config.configFile)

	// Start file monitoring
	go NewFileMonitor(config, logs)

	// Start a TCP socket server
	go NewWebSocketServer(config, logs)

	// Start HTTP server
	go NewHTTPServer()

	// Block
	select {}
}
