package main

import (
	"io"
	"log"
	"net/http"
	"time"

	_ "net/http/pprof"

	"github.com/miku/gdate"
)

func handler(w http.ResponseWriter, req *http.Request) {
	s := gdate.Preamble() + time.Now().Format(time.RFC3339) + "\n"
	if _, err := io.WriteString(w, s); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
}

func main() {
	http.HandleFunc("/", handler)
	addr := "0.0.0.0:8000"
	log.Println("http://" + addr)
	log.Fatal(http.ListenAndServe(addr, nil))
}
