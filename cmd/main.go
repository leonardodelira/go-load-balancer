package main

import (
	"fmt"
	"net/http"

	"github.com/leonardodelira/go-load-balancer/cmd/core"
)

func main() {
	servers := []core.Server{
		core.NewServer("https://www.google.com/"),
		core.NewServer("https://www.bing.com"),
		core.NewServer("https://www.duckduckgo.com"),
	}

	lb := core.NewLoadBalancer("8080", servers)
	handleRedirect := func(rw http.ResponseWriter, req *http.Request) {
		lb.ServeProxy(rw, req)
	}

	http.HandleFunc("/", handleRedirect)
	fmt.Printf("serving requests at 'localhost:%s'\n", lb.Port)
	http.ListenAndServe(":"+lb.Port, nil)
}
