package wstonats

import (
	"github.com/gorilla/websocket"
	"strings"
	"testing"
	"time"
	"sync"
	"fmt"
)

type logWriter struct{}

func (w *logWriter) Write(msg []byte) (n int, err error) {
	return fmt.Println(string(msg))
}

func TestGateWay(t *testing.T) {
	proxySetting := new(ProxySetting)
	proxySetting.WsAddr = "0.0.0.0:8910"
	proxySetting.NatsAddr = "0.0.0.0:4222"
	proxySetting.DebugLevel = DebugInfo
	proxySetting.LogWriter = new(logWriter)
	go Start(proxySetting)

	time.Sleep(100 * time.Millisecond)
	wsCon, _, err := websocket.DefaultDialer.Dial("ws://0.0.0.0:8910", nil)
	if err != nil {
		t.Error(err)
	}
	defer wsCon.Close()

	wg := sync.WaitGroup{}
	wg.Add(1)
	timer := time.NewTimer(3 * time.Second)
	go func() {
		go func() {
			<-timer.C
			t.Error("Timout come")
		}()
		for {
			_, message, err := wsCon.ReadMessage()
			fmt.Println(`=====FROM WS: ` + string(message))
			if err != nil {
				t.Error(err)
				wg.Done()
				return
			}
			if strings.Contains(string(message), "Hello test") {
				wg.Done()
				return
			}
		}
	}()

	err = wsCon.WriteMessage(websocket.TextMessage, []byte("CONNECT {\"verbose\":true,\"pedantic\":false,\"tls_required\":false,\"name\":\"\",\"lang\":\"go\",\"version\":\"1.2.2\",\"protocol\":1}\r\n"))
	err = wsCon.WriteMessage(websocket.TextMessage, []byte("sub foo.* 1\r\n"))
	err = wsCon.WriteMessage(websocket.TextMessage, []byte("pub foo.bar 10\r\nHello test\r\n"))
	if err != nil {
		t.Error(err)
	}
	wg.Wait()
	if !timer.Stop() {
		<-timer.C
	}
}
