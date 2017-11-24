# Example

1. Run NATS Server at first. You can use a Docker:

```bash
docker run -d --name nats-main -p 4222:4222 -p 6222:6222 -p 8222:8222 nats
```

2. Clone this repo and run:

```bash
go run ./example/gw.go
```

3. Just open "page-example.html" in the browser and look to browser dev-console.