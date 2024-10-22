package main

import (
	"fmt"
	"log"

	"github.com/YamatoKato/st-go-web-dev/app/server"
)

const (
	host = "0.0.0.0" // すべてのネットワークインターフェースをlistenする
	port = "8080"    // セキュリティグループ等で指定したポート番号と合わせる
)

func main() {
	s := server.New()
	err := s.Start(fmt.Sprintf("%s:%s", host, port))
	if err != nil {
		s.Stop()
		log.Fatalf("server stopped with error: %s", err)
	}
}
