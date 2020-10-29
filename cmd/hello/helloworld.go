package main

import (
	"flag"
  "fmt"
	"net/http"

	log "github.com/sirupsen/logrus"
)

const msg = "Hello, world!"

func handler(w http.ResponseWriter, r *http.Request) {
  fmt.Fprintf(w, msg)
}

func main() {
	var (
		addr  = flag.String("addr", ":8080", "listen address and port")
		httpd = flag.Bool("httpd", false, "enable http server")
	)

	flag.Parse()

	if *httpd {
		log.Printf("start %s on %s", "hello-world", *addr)
		http.HandleFunc("/", handler)
		if err := http.ListenAndServe(*addr, nil); err != nil {
			log.Fatalf("%v", err)
		}
	} else {
    fmt.Println(msg)
  }
}
