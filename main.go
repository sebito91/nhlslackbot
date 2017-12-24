// +build linux

package main

import (
	"context"
	"log"
	"net/http"
	"os"

	"github.com/sebito91/nhlslackbot/post"

	_ "expvar"
	_ "net/http/pprof"
)

func main() {
	lg := log.New(os.Stdout, "slack-bot: ", log.Lshortfile|log.LstdFlags)
	ctx := context.Background()

	s, err := post.New()
	if err != nil {
		lg.Fatal(err)
	}
	s.Logger = lg

	if err := s.Run(ctx); err != nil {
		s.Logger.Fatal(err)
	}

	handler, err := s.NewHandler()
	if err != nil {
		s.Logger.Fatal(err)
	}

	lg.Fatal(http.ListenAndServe(":9191", handler))
}
