# NATS <=> WebSocket GateWay

Easy way to use NATS EventBus directly on front-end side

## Features
- High performance proxy server
- TLS support
- Protect data transfer by using own firewall

## How it works

Use on your front side js library from [here](https://github.com/isobit/websocket-nats)
See example [here](https://github.com/shubinmi/wstonats/tree/master/example)

## How to implement

```go
package main

import (
	"github.com/shubinmi/wstonats"
)

func main() {
	proxySetting := new(wstonats.ProxySetting)
	
	proxySetting.WsAddr = "0.0.0.0:8910"
	proxySetting.NatsAddr = "0.0.0.0:4222"
	proxySetting.DebugLevel = wstonats.DebugInfo
    
	wstonats.Start(proxySetting)
}
```
