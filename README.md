# drekle/dns
A DNS server made for machines, not people.  DNS updates can event driven, but they require an API to do this.  The goal here is to use GRPC to create client libraries in all languages including a HTTP/JSON Rest implementation.

## My use case
On the creation of an ingress resource in kubernetes you might want to notify a DNS server to create an A record so that traffic could be routed to ingress controllers within the cluster from inside your organization.  A conditional forwarder can be used to route to the dns server provided here, and a separate service ( which I will also provide ) will update this dns server based off of ingresses CRUD events within the cluster.

## Docker
The build container will be available at drekle/dns.

This makes use heavily on the dns library github.com/miekg/dns.  See this page for reference.  

This DNS server listens on a couple of ports by default.
53 TCP - DNS
53 UDP - DNS
80 - The port for http rest requests
50051 - GRPC port (The HTTP requests get forwarded to this client to be handled.)

### Reasoning
The kubernetes control plane is event driven.  Because of this I need a DNS server that I can programatically update based off of these events.  This DNS server
uses gRPC to provide CRUD endpoints for updating DNS entries.

## Run step
This application is intended to be ran outside of kubernetes so that it can be bound to TCP 53 and UDP 53.

Below is an example cmd to run the docker container...

```
docker build -t dnsserver .
docker run -itd --name dnsserver --restart always --net=host dnsserver -database <MONGODB_ADDR>
```

### BUILD
Generate the required files

```
protoc api.proto -IC:\Users\dlemon\go\src\github.com\googleapis\googleapis -I. --go_out=plugins=grpc:lib/go/v1
protoc api.proto -IC:\Users\dlemon\go\src\github.com\googleapis\googleapis -I. --grpc-gateway_out=lib/go/v1
protoc api.proto -I. -IC:\Users\dlemon\go\src\github.com\googleapis\googleapis --swagger_out=swagger
go generate .
```
We build client libs only for golang, add languages here if desirable.  Recommend doing this even if using http/json rest requests as you can marshal the arguement structs.

Run the swagger gen utility

```
cd swagger
go run swaggergen.go
```

Build the Binary for the os of choice.  Linux used below...
```
set GOOS=linux
set GOARCH=amd64
go build .
```

Build the docker container
```
docker build -t dns .
```

Below an example of our run
```
docker run -itd --restart always --name dns --net=host dns -database mongo
```
Instead of specifying the database, you could also add mongo to your hosts file.

