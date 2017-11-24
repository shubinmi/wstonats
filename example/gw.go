package main

import (
	"fmt"
	"github.com/shubinmi/wstonats"
	"net/http"
)

type logWriter struct{}

func (w *logWriter) Write(msg []byte) (n int, err error) {
	return fmt.Println(string(msg))
}

type firewall struct{}

func (f *firewall) Allow(msg []byte, requestHeader http.Header) bool {
	return true
}

func main() {
	proxySetting := new(wstonats.ProxySetting)
	proxySetting.WsAddr = "0.0.0.0:8910"
	proxySetting.NatsAddr = "0.0.0.0:4222"
	proxySetting.DebugLevel = wstonats.DebugInfo
	proxySetting.LogWriter = new(logWriter)
	proxySetting.Firewall = new(firewall)

	wstonats.Start(proxySetting)
}
