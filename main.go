package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"os"

	"github.com/YamatoKato/st-go-web-dev/config"
)

func main() {
	if err := run(context.Background()); err != nil {
		log.Printf("run() err: %v", err)
		os.Exit(1)
	}
}

func run(ctx context.Context) error {
	cfg, err := config.New()
	if err != nil {
		return err
	}
	l, err := net.Listen("tcp", fmt.Sprintf(":%d", cfg.Port))
	if err != nil {
		log.Fatalf("net.Listen() err: %v", err)
	}
	url := fmt.Sprintf("http://%s", l.Addr().String())
	log.Printf("server listening at %s", url)
	mux := NewMux()
	s := NewServer(l, mux)

	return s.Run(ctx)
}
