package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"mime"
	"net"
	"net/http"
	"strconv"
	"strings"

	"github.com/drekle/dns-server/pkg/ui/data/swagger"

	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"github.com/philips/go-bindata-assetfs"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	pb "github.com/drekle/dns-server/pkg/lib/go/v1"
	"github.com/miekg/dns"
)

// grpcHandlerFunc returns an http.Handler that delegates to grpcServer on incoming gRPC
// connections or otherHandler otherwise. Copied from cockroachdb.
func grpcHandlerFunc(grpcServer *grpc.Server, otherHandler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.ProtoMajor == 2 && strings.Contains(r.Header.Get("Content-Type"), "application/grpc") {
			grpcServer.ServeHTTP(w, r)
		} else {
			otherHandler.ServeHTTP(w, r)
		}
	})
}

func serveSwagger(mux *http.ServeMux) {
	mux.HandleFunc("/swagger.json", func(w http.ResponseWriter, req *http.Request) {
		io.Copy(w, strings.NewReader(pb.Swagger))
	})
	mime.AddExtensionType(".svg", "image/svg+xml")

	// Expose files in third_party/swagger-ui/ on <host>/swagger-ui
	fileServer := http.FileServer(&assetfs.AssetFS{
		Asset:    swagger.Asset,
		AssetDir: swagger.AssetDir,
		Prefix:   "third_party/swagger-ui",
	})
	prefix := "/ui/"
	mux.Handle(prefix, http.StripPrefix(prefix, fileServer))
}

func startDNSServer(dnsServer *server) {
	go func() {
		udpSrv := &dns.Server{Addr: ":" + strconv.Itoa(53), Net: "udp"}
		udpSrv.Handler = dnsServer
		if err := udpSrv.ListenAndServe(); err != nil {
			log.Fatalf("Failed to set udp listener %s\n", err.Error())
		}
	}()
	//Listen TCP
	go func() {
		tcpSrv := &dns.Server{Addr: ":" + strconv.Itoa(53), Net: "tcp"}
		tcpSrv.Handler = dnsServer
		if err := tcpSrv.ListenAndServe(); err != nil {
			log.Fatalf("Failed to set tcp listener %s\n", err.Error())
		}
	}()
}

func startGRPCServer(port string, dnsServer *server) *grpc.Server {
	s := grpc.NewServer()
	go func() {
		pb.RegisterDNSServiceServer(s, dnsServer)
		reflection.Register(s)
		lis, err := net.Listen("tcp", port)
		if err != nil {
			log.Fatalf("failed to listen: %v", err)
		}
		log.Printf("Serving gRPC on %s", port)
		if err := s.Serve(lis); err != nil {
			log.Fatalf("failed to serve: %v", err)
		}
	}()
	return s
}

func main() {
	grpcPort := flag.String("port", ":50051", "port which grpc listens.")
	database := flag.String("database", "mongodb", "port which grpc listens.")
	flag.Parse()

	//Listen on 53 TCP/UDP for dns forever
	dnsServer := NewServer(*database)
	startDNSServer(dnsServer)
	//Listen UDP

	//gRPC
	s := startGRPCServer(*grpcPort, dnsServer)

	//Overarching HTTP Mux
	mux := http.NewServeMux()

	//Setup gRPC gateway to support HTTP REST -> generated methods
	gwmux := runtime.NewServeMux()
	err := pb.RegisterDNSServiceHandlerFromEndpoint(context.Background(), gwmux, *grpcPort, []grpc.DialOption{grpc.WithInsecure()})
	if err != nil {
		fmt.Printf("serve: %v\n", err)
		return
	}
	mux.Handle("/", gwmux)

	//Create swagger endpoints for visualization
	serveSwagger(mux)

	//Serve http
	conn, err := net.Listen("tcp", fmt.Sprintf(":%d", 80))
	if err != nil {
		panic(err)
	}
	srv := &http.Server{
		Addr:    ":80",
		Handler: grpcHandlerFunc(s, mux),
	}
	err = srv.Serve(conn)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}

	return
}
