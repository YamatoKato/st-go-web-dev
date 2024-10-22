package server

import (
	"log"
	"net/http"
	"sync/atomic"
	"time"

	"github.com/gin-gonic/gin"
)

type Server struct {
	counter int64

	server *http.Server
	router *gin.Engine
}

func New() *Server {
	router := gin.Default()
	server := &Server{
		router:  router,
		counter: int64(0),
	}

	router.GET("/", server.CounterHandler)
	router.GET("/health_check", server.HealthCheckHandler)

	return server
}

func (s *Server) HealthCheckHandler(ctx *gin.Context) {
	ctx.JSON(200, gin.H{"success": true})
}

func (s *Server) CounterHandler(ctx *gin.Context) {
	// atomic.AddInt64は、複数のgoroutineから同時に呼び出されても安全にカウンタをインクリメントする
	counter := atomic.AddInt64(&s.counter, 1)
	ctx.JSON(200, gin.H{"counter": counter})
}

func (s *Server) Start(address string) error {
	s.server = &http.Server{
		Addr:        address,          // e.g. ":8080"
		Handler:     s.router,         // gin.Engineをhttp.Handlerとして渡す
		ReadTimeout: 10 * time.Second, // タイムアウト設定
	}
	log.Printf("start server on %s", address)

	return s.server.ListenAndServe()
}

func (s *Server) Stop() error {
	log.Printf("stop server")
	if s.server != nil {
		return s.server.Close()
	}

	return nil
}
