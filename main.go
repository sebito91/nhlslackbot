// +build linux

package main

import (
	"fmt"
	"log"

	"github.com/sebito91/nhlslackbot/fetch"
)

func main() {
	fmt.Printf("this is a fun project\n")

	n := fetch.New()

	if err := n.GetSchedule(); err != nil {
		log.Fatalf("error in fetch schedule: %+v", err)
	}
}
