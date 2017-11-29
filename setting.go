package wstonats

import (
	"crypto/tls"
	"io"
	"net/http"
)

//MsgFirewall noinspection GoUnusedExportedInterface
type MsgFirewall interface {
	Allow(msg *[]byte, requestHeader http.Header) bool
}

//ProxySetting noinspection GoUnusedExportedStruct
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

//DebugNon noinspection GoUnusedExportedConst
const (
	DebugNon  = 0
	DebugErr  = 1
	DebugInfo = 2
)