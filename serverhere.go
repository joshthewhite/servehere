package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
)

var (
	port int
)

func init() {
	flag.IntVar(&port, "port", 8888, "the port on which the server will listen")
}

func main() {
	flag.Parse()

	cwd, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	addr := fmt.Sprintf(":%d", port)
	fmt.Println("Listening on", addr)

	panic(http.ListenAndServe(addr, addDefaultHeaders(http.FileServer(http.Dir(cwd)))))
}

func addDefaultHeaders(h http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        w.Header().Set("Access-Control-Allow-Origin", "*")
        h.ServeHTTP(w, r)
    })
}

