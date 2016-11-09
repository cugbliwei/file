package main

import (
	"flag"
	"fmt"
	"log"
	"net"
	"net/http"
	_ "net/http/pprof"
)

func HandleHealth(w http.ResponseWriter, req *http.Request) {
	fmt.Fprint(w, "ok")
}

func main() {
	port := flag.String("port", "8001", "port number")
	flag.Parse()

	http.HandleFunc("/health", HandleHealth)
	http.Handle("/temp/", http.StripPrefix("/temp/", http.FileServer(http.Dir("./temp"))))
	l, e := net.Listen("tcp", ":"+*port)
	if e != nil {
		log.Fatal("listen error:", e)
	}
	http.Serve(l, nil)
}
