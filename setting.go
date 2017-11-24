package wstonats

import "io"

type MsgFirewall interface {
	Allow(msg []byte) bool
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
