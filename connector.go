package wstonats

import (
	"bufio"
	"crypto/tls"
	"fmt"
	"github.com/gobwas/ws"
	"github.com/gobwas/ws/wsutil"
	"net"
	"net/http"
	"time"
)

type connector struct {
	id          int
	nats        net.Conn
	ws          net.Conn
	toWs        chan []byte
	toNats      chan []byte
	natsScanner *bufio.Scanner
	reqHeader   http.Header
}

func newConnector(w http.ResponseWriter, r *http.Request) *connector {
	rHeader := r.Header
	rHeader.Add("Remote-Address", r.RemoteAddr)

	r.UserAgent()
	wsConn, _, _, err := ws.UpgradeHTTP(r, w)
	if err != nil {
		writeLog([]byte("Cannot upgrade HTTP to WS errText:"+err.Error()), DebugErr)
		panic(err)
	}

	natsConnInit := func(useTls bool) net.Conn {
		switch useTls {
		case true:
			conn, err := tls.Dial("tcp", proxySetting.NatsAddr, proxySetting.NatsTlsConf)
			if err != nil {
				writeLog([]byte("Cannot connect to NATS errText:"+err.Error()), DebugErr)
				panic(err)
			}
			return conn
		default:
			conn, err := net.Dial("tcp", proxySetting.NatsAddr)
			if err != nil {
				writeLog([]byte("Cannot connect to NATS errText:"+err.Error()), DebugErr)
				panic(err)
			}
			return conn
		}
	}
	natsConn := natsConnInit(proxySetting.NatsTls)

	if err != nil {
		writeLog([]byte("Cannot connect to NATS errText:"+err.Error()), DebugErr)
		panic(err)
	}
	id++
	return &connector{
		id:          id,
		nats:        natsConn,
		ws:          wsConn,
		toWs:        make(chan []byte),
		toNats:      make(chan []byte),
		natsScanner: bufio.NewScanner(bufio.NewReader(natsConn)),
		reqHeader:   rHeader,
	}
}

func (c *connector) clear() {
	writeLog([]byte(fmt.Sprintf("Connector closed id:%v", c.id)), DebugInfo)
	if c.ws != nil {
		_ = c.ws.Close()
	}
	if c.nats != nil {
		_ = c.nats.Close()
	}
	c.ws = nil
	c.nats = nil
	c.toWs = nil
	c.toNats = nil
}

func (c *connector) pullFromWs() {
	defer c.clear()
	for {
		if c.ws == nil {
			return
		}
		msg, op, err := wsutil.ReadClientData(c.ws)
		if err != nil {
			return
		}
		writeLog([]byte(fmt.Sprintf("MSG from WS for id:%v with opCode:%v", c.id, op)), DebugInfo)
		if op == ws.OpClose || op == ws.OpContinuation {
			return
		}
		if op != ws.OpText {
			continue
		}
		if proxySetting.Firewall != nil && !proxySetting.Firewall.Allow(&msg, c.reqHeader) {
			c.toWs <- []byte("-ERR 'Invalid Subject'\r\n")
			writeLog([]byte(fmt.Sprintf("MSG not allowed from WS for id:%v with msg:%v", c.id, msg)), DebugInfo)
			time.Sleep(1 * time.Second)
			return
		}
		writeLog([]byte(fmt.Sprintf("MSG from WS for id:%v with msg:%s", c.id, msg)), DebugInfo)
		c.toNats <- msg
	}
}

func (c *connector) pushToNats() {
	for {
		if c.nats == nil {
			return
		}
		msg := <-c.toNats
		writeLog([]byte(fmt.Sprintf("MSG to NATS for id:%v with msg:%s", c.id, msg)), DebugInfo)
		_, _ = c.nats.Write(msg)
	}
}

func (c *connector) pushToWs() {
	for {
		if c.ws == nil {
			return
		}
		msg := <-c.toWs
		writeLog([]byte(fmt.Sprintf("MSG to WS for id:%v with msg:%s", c.id, msg)), DebugInfo)
		err := wsutil.WriteServerMessage(c.ws, ws.OpText, msg)
		if err != nil {
			c.clear()
			return
		}
	}
}

func (c *connector) pullFromNats() {
	for c.natsScanner.Scan() {
		if c.nats == nil {
			break
		}
		msg := c.natsScanner.Bytes()
		if len(msg) == 0 {
			continue
		}
		writeLog([]byte(fmt.Sprintf("MSG from NATS for id:%v with msg:%s", c.id, msg)), DebugInfo)
		c.toWs <- append(msg, '\r', '\n')
	}
	if err := c.natsScanner.Err(); err != nil {
		c.clear()
	}
	c.natsScanner = nil
}

func (c *connector) init() {
	go c.pullFromWs()
	go c.pushToNats()
	go c.pullFromNats()
	go c.pushToWs()
}
