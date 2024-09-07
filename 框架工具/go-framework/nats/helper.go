package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/nats-io/nats.go"
)

func connect(addr string, opts ...nats.Option) *nats.Conn {
	nc, err := nats.Connect(addr, opts...)
	if err != nil {
		log.Fatal(err)
	}

	return nc
}

func blocking() {
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM)
	<-sig
}
