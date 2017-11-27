package wstonats

import (
	"io"
	"net/http"
	"crypto/tls"
)

type MsgFirewall interface {
	Allow(msg *[]byte, requestHeader http.Header) bool
}

type ProxySetting struct {
	WsAddr      string
	WsTls       bool
	WsTlsCert   string
	WsTlsKey    string
	NatsAddr    string
	NatsTls     bool
	NatsTlsConf *tls.Config
	DebugLevel  int
	LogWriter   io.Writer
	Firewall    MsgFirewall
}

const (
	DebugNon  = 0
	DebugErr  = 1
	DebugInfo = 2
)
