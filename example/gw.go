package main

import (
	"fmt"
)

type logWriter struct{}

func (w *logWriter) Write(msg []byte) (n int, err error) {
	return fmt.Println(string(msg))
}

type firewall struct{}

func (f *firewall) Allow(msg []byte) bool {
	return true
}

func main() {
	//proxySetting.WsAddr = "0.0.0.0:8910"
	//proxySetting.NatsAddr = "0.0.0.0:4222"
	//proxySetting.DebugLevel = DebugInfo
	//proxySetting.LogWriter = new(logWriter)
	//proxySetting.Firewall = new(firewall)
}
