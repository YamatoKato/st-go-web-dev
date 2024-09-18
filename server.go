package main

import (
	"context"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"golang.org/x/sync/errgroup"
)

type Server struct {
	srv *http.Server
	l   net.Listener
}

func NewServer(l net.Listener, mux http.Handler) *Server {
	return &Server{
		srv: &http.Server{
			Handler: mux,
		},
		l: l,
	}
}

func (s *Server) Run(ctx context.Context) error {
	// グレースフルシャットダウンのためのコンテキストを作成
	ctx, stop := signal.NotifyContext(ctx, os.Interrupt, syscall.SIGTERM)
	// シグナルを受け取ったらstop()を呼び出す
	defer stop()

	eg, ctx := errgroup.WithContext(ctx)
	// 別のgoroutineでサーバーを起動
	eg.Go(func() error {
		if err := s.srv.Serve(s.l); err != nil && err != http.ErrServerClosed {
			return err
		}
		return nil
	})

	<-ctx.Done() // ここで終了シグナルを待つ. Done()は親コンテキストの終了を検知する
	if err := s.srv.Shutdown(ctx); err != nil {
		return err
	}
	return eg.Wait()
}
