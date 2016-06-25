package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

/* understand the type Server struct */
/*

type Server struct {
	Addr	string
	Handler Handler
	ReadTimeout	time.Duration
	WriteTimeout	time.Duration
	MaxHeaderBytes	int
	TLSConfig	*tls.Config
	TLSNextProto	map[string]func(*Server, *tls.Conn, Handler)
	ConnState	func(net.Conn, ConnState)
	ErrorLog	*log.Logger
}

*/

func gwpHandler(msg string) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, msg)
	})
}

func shortHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "contact me at ardeshir.org at gmail.com")
}

func main() {
	// mux := http.NewServeMux()

	http.Handle("/", gwpHandler("Hi from gwp! Try /about"))
	http.Handle("/about", gwpHandler("About gwp.."))
	http.HandleFunc("/contact", shortHandler)

	log.Println("We're up on 9090...")

	server := &http.Server{

		Addr:            ":9090",
		ReadTimeout:     10 * time.Second,
		WriteTimeout:    10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	server.ListenAndServe()
}
