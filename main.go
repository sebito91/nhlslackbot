// +build linux

package main

import (
	"log"

	"github.com/sebito91/nhlslackbot/fetch"
)

func main() {
	n := fetch.New()

	if err := n.GetSchedule(); err != nil {
		log.Fatalf("error in fetch schedule: %+v", err)
	}
}
