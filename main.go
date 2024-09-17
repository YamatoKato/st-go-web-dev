package main

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"golang.org/x/sync/errgroup"
)

func main() {
	//
}

func run(ctx context.Context) error {
	s := &http.Server{
		Addr: ":18080",
		Handler: http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			fmt.Fprintf(w, "Hello, %s", r.URL.Path[1:])
		}),
	}

	eg, ctx := errgroup.WithContext(ctx)
	// 別のgoroutineでサーバーを起動
	eg.Go(func() error {
		if err := s.ListenAndServe(); err != nil && err != http.ErrServerClosed {
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
