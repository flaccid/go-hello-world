package main

import (
        "flag"
        "fmt"
        "github.com/gorilla/handlers"
        log "github.com/sirupsen/logrus"
        "net/http"
        "os"
)

var (
        addr  = flag.String("addr", ":8080", "listen address and port")
        msg   = flag.String("msg", "Hello, world!", "the greeting to echo")
        httpd = flag.Bool("httpd", false, "enable http server")
)

func handler(w http.ResponseWriter, r *http.Request) {
        fmt.Fprintf(w, *msg)
}

func main() {
        flag.Parse()

        if *httpd {
                log.Printf("start %s on %s", "hello-world", *addr)
                r := http.NewServeMux()
                r.Handle("/", handlers.LoggingHandler(os.Stdout, http.HandlerFunc(handler)))
                http.ListenAndServe(*addr, handlers.CompressHandler(r))
        } else {
                fmt.Println(*msg)
        }

}
