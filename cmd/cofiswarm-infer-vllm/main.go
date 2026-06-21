package main

import (
	"flag"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/keepdevops/cofiswarm-infer-vllm/internal/bus"
	"github.com/keepdevops/cofiswarm-observer-sdk/pkg/servicecomponent"
)

func main() {
	addr := flag.String("listen", ":8088", "health/metadata port (HTTP mode)")
	busMode := flag.Bool("bus", false, "announce + serve .infer.vllm.* on the NATS observer bus instead of HTTP")
	natsURL := flag.String("nats", "nats://127.0.0.1:4222", "NATS URL (bus mode)")
	flag.Parse()

	if *busMode {
		serveBus(*natsURL)
		return
	}

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

// serveBus announces infer-vllm on the observer bus and serves its .infer.vllm.* capability
// subjects until SIGINT/SIGTERM, when it says goodbye so presence flips offline cleanly.
func serveBus(url string) {
	nc, err := servicecomponent.Connect(url, "cofiswarm-infer-vllm")
	if err != nil {
		log.Fatalf("bus connect %s: %v", url, err)
	}
	defer nc.Close()
	comp := servicecomponent.New(nc, "infer-vllm", "infer-vllm", bus.Routes("vllm"))
	if err := comp.Start(); err != nil {
		log.Fatalf("bus start: %v", err)
	}
	defer comp.Shutdown()
	log.Printf("infer-vllm on bus %s (.infer.vllm.info/.health)", url)

	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM)
	<-sig
	log.Print("infer-vllm bus stopping")
}
