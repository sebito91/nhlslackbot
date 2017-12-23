// +build linux

package main

import (
	"log"

	"github.com/sebito91/nhlslackbot/fetch"
	"github.com/sebito91/nhlslackbot/post"
	"go.uber.org/zap"

	_ "expvar"
	"net/http"
	_ "net/http/pprof"
)

func main() {
	logger, err := zap.NewProduction()
	if err != nil {
		log.Fatal("could not instantiate logger")
	}
	defer logger.Sync()
	lg := logger.Sugar()

	n := fetch.New()

	if err := n.GetSchedule(); err != nil {
		lg.Fatalf("error in fetch schedule: %+v", err)
	}

	lg.Infof("got the schedule, doneskii -- %s", "yoyo")

	//	ctx := context.Background()

	s, err := post.New()
	if err != nil {
		lg.Fatal(err)
	}
	s.Logger = lg

	//	var wg sync.WaitGroup
	//	wg.Add(1)
	//
	//	s := post.New()
	//	go s.Run()
	//	wg.Wait()

	lg.Fatal(http.ListenAndServe(":9191", nil))
}
