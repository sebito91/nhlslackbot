// +build linux

package main

import (
	"context"
	"log"
	"os"

	"github.com/sebito91/nhlslackbot/fetch"
	"github.com/sebito91/nhlslackbot/post"

	_ "expvar"
	"net/http"
	_ "net/http/pprof"
)

func main() {
	lg := log.New(os.Stdout, "slack-bot: ", log.Lshortfile|log.LstdFlags)
	n := fetch.New()

	if err := n.GetSchedule(); err != nil {
		lg.Fatalf("[ERROR] error in fetch schedule: %+v", err)
	}

	lg.Printf("[DEBUG] got the schedule, doneskii -- %s", "yoyo")

	ctx := context.Background()

	s, err := post.New()
	if err != nil {
		lg.Fatal(err)
	}
	s.Logger = lg

	if err := s.Run(ctx); err != nil {
		lg.Fatal(err)
	}

	http.HandleFunc("/", post.IndexHandler)
	lg.Fatal(http.ListenAndServe(":9191", nil))
}
