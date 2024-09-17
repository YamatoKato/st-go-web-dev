package main

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"testing"

	"golang.org/x/sync/errgroup"
)

func TestRun(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	eg, ctx := errgroup.WithContext(ctx)
	eg.Go(func() error {
		return run(ctx)
	})
	in := "message"
	rsp, err := http.Get("http://localhost:18080/" + in)
	if err != nil {
		t.Errorf("http.Get() err: %v", err)
	}
	defer rsp.Body.Close()

	got, err := io.ReadAll(rsp.Body)
	if err != nil {
		t.Fatalf("io.ReadAll() err: %v", err)
	}

	// HTTPサーバーの戻り値を検証する
	want := fmt.Sprintf("Hello, %s", in)
	if string(got) != want {
		t.Errorf("http.Get() = %s; want %s", got, want)
	}

	// run()の終了を送信する
	cancel()

	// run()のもどり値を検証する
	if err := eg.Wait(); err != nil {
		t.Fatal(err)
	}
}
