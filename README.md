# NATS <-> WebSocket gateway

Easy way to use all NATS EventBus directly on front-end side

## How to use

Fetch the source:

```bash
go get -u github.com/orus-io/nats-websocket-gw
```

Install and run the default binary

```bash
go install github.com/orus-io/nats-websocket-gw/cmd/nats-websocket-gw
nats-websocket-gw --no-origin-check
```

and/or integrate it in your http server:

```go
package main

import (
	"net/http"

	"github.com/orus-io/nats-websocket-gw"
)

func main() {
	gateway := gw.NewGateway(gw.Settings{
		NatsAddr: "localhost:4222",
	})
	http.HandleFunc("/nats", gateway.Handler)
	http.ListenAndServe("0.0.0.0:8910", nil)
}
```
