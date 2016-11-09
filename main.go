package main

import (
	"flag"
	"log"
	"net"
	"net/http"
	_ "net/http/pprof"
)

func main() {
	port := flag.String("port", "8001", "port number")
	flag.Parse()
	http.Handle("/temp/", http.StripPrefix("/site/", http.FileServer(http.Dir("./site"))))
	l, e := net.Listen("tcp", ":"+*port)
	if e != nil {
		log.Fatal("listen error:", e)
	}
	http.Serve(l, nil)
}
