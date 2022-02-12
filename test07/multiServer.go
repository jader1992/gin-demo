package main

import (
	"github.com/gin-gonic/gin"
	"golang.org/x/sync/errgroup"
	"log"
	"net/http"
	"time"
)

var (
	g errgroup.Group
)

func route01() http.Handler  {
	e := gin.New()
	e.Use(gin.Recovery())
	e.GET("/", func(context *gin.Context) {
		context.JSON(http.StatusOK, gin.H{
			"code": http.StatusOK,
			"error": "welcome server 01",
		})
	})
	return e
}

func route02() http.Handler  {
	e := gin.New()
	e.Use(gin.Recovery())
	e.GET("/", func(context *gin.Context) {
		context.JSON(http.StatusOK, gin.H{
			"code": http.StatusOK,
			"error": "welcome server 02",
		})
	})
	return e
}

func main() {
	server01 := &http.Server{
		Addr:              ":8080",
		Handler:           route01(),
		TLSConfig:         nil,
		ReadTimeout:       5 * time.Second,
		ReadHeaderTimeout: 10 * time.Second,
		WriteTimeout:      0,
		IdleTimeout:       0,
		MaxHeaderBytes:    0,
		TLSNextProto:      nil,
		ConnState:         nil,
		ErrorLog:          nil,
		BaseContext:       nil,
		ConnContext:       nil,
	}

	server02 := &http.Server{
		Addr:              ":8081",
		Handler:           route02(),
		TLSConfig:         nil,
		ReadTimeout:       5 * time.Second,
		ReadHeaderTimeout: 10 * time.Second,
		WriteTimeout:      0,
		IdleTimeout:       0,
		MaxHeaderBytes:    0,
		TLSNextProto:      nil,
		ConnState:         nil,
		ErrorLog:          nil,
		BaseContext:       nil,
		ConnContext:       nil,
	}

	g.Go(func() error {
		return server01.ListenAndServe()
	})

	g.Go(func() error {
		return server02.ListenAndServe()
	})

	if err := g.Wait(); err != nil {
		log.Fatal(err)
	}
}