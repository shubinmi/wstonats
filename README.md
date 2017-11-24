# NATS <=> WebSocket GateWay

Easy way to use NATS EventBus directly on front-end side

## How to use

```go
package main

import (
	"net/http"
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

Or see to example folder
