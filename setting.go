package wstonats

import (
	"io"
	"net/http"
)

type MsgFirewall interface {
	Allow(msg []byte, requestHeader http.Header) bool
}

type ProxySetting struct {
	WsAddr     string
	NatsAddr   string
	DebugLevel int
	LogWriter  io.Writer
	Firewall   MsgFirewall
}

const (
	DebugNon  = 0
	DebugErr  = 1
	DebugInfo = 2
)
