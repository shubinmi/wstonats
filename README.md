# NATS <=> WebSocket GateWay

Easy way to use NATS EventBus directly on front-end side

[![Open Source Love](https://badges.frapsoft.com/os/v2/open-source.svg?v=103)](https://github.com/shubinmi/salesforce-bulk-api) [![MIT Licence](https://badges.frapsoft.com/os/mit/mit.svg?v=103)](https://opensource.org/licenses/mit-license.php) 

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
