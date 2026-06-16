package main

import (
	"flag"
	"log"
	"net/http"
)

func main() {
	addr := flag.String("listen", ":8088", "health/metadata port")
	flag.Parse()
	mux := http.NewServeMux()
	mux.HandleFunc("/healthz", func(w http.ResponseWriter, _ *http.Request) {
		w.Write([]byte("ok"))
	})
	mux.HandleFunc("/v1/info", func(w http.ResponseWriter, _ *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write([]byte(`{"engine":"vllm","stub":true,"note":"run deploy/Dockerfile for full vllm-metal"}`))
	})
	log.Printf("infer-vllm metadata on %s", *addr)
	log.Fatal(http.ListenAndServe(*addr, mux))
}
