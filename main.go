package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"net/http"
	"os"

	"golang.org/x/sync/errgroup"
)

func main() {
	if len(os.Args) != 2 {
		log.Printf("Usage: %s <name>", os.Args[0])
		os.Exit(1)
	}

	p := os.Args[1]
	l, err := net.Listen("tcp", ":"+p)
	if err != nil {
		log.Fatalf("net.Listen() err: %v", err)
	}
	if err := run(context.Background(), l); err != nil {
		log.Printf("failed to run: %v", err)
		os.Exit(1)
	}
}

func run(ctx context.Context, l net.Listener) error {
	s := &http.Server{
		Handler: http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			fmt.Fprintf(w, "Hello, %s", r.URL.Path[1:])
		}),
	}

	eg, ctx := errgroup.WithContext(ctx)
	// 別のgoroutineでサーバーを起動
	eg.Go(func() error {
		if err := s.Serve(l); err != nil && err != http.ErrServerClosed {
			log.Printf("failed to listen and server: %v", err)
			return err
		}
		return nil
	})

	<-ctx.Done() // ここで終了シグナルを待つ. Done()は親コンテキストの終了を検知する
	if err := s.Shutdown(ctx); err != nil {
		log.Printf("failed to shutdown server: %v", err)
	}

	return eg.Wait() // すべてのgoroutineが終了するまで待つ
}
