package wstonats

import (
	"testing"
	"github.com/gorilla/websocket"
	"github.com/shubinmi/wstonats"
	"strings"
)

func TestGateWay(t *testing.T) {
	proxySetting := new(wstonats.ProxySetting)
	proxySetting.WsAddr = "0.0.0.0:8910"
	proxySetting.NatsAddr = "0.0.0.0:4222"
	proxySetting.DebugLevel = wstonats.DebugInfo
	go wstonats.Start(proxySetting)

	wsCon, _, err := websocket.DefaultDialer.Dial("ws://0.0.0.0:8910", nil)
	if err != nil {
		t.Error(err)
	}
	defer wsCon.Close()

	go func() {
		for {
			_, message, err := wsCon.ReadMessage()
			if err != nil {
				t.Error(err)
				return
			}
			if strings.Contains(string(message), "hello test") {
				return
			}
		}
	}()

	err = wsCon.WriteMessage(websocket.TextMessage, []byte("[CONNECT {\"verbose\":false,\"pedantic\":false,\"tls_required\":false,\"name\":\"\",\"lang\":\"go\",\"version\":\"1.2.2\",\"protocol\":1}]\r\n"))
	err = wsCon.WriteMessage(websocket.TextMessage, []byte("SUB foo 1\r\n"))
	err = wsCon.WriteMessage(websocket.TextMessage, []byte("PUB foo 10\r\nhello test\r\n"))
	if err != nil {
		t.Error(err)
	}
}
