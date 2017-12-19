// +build linux

package main

import (
	"fmt"
	"log"

	"github.com/sebito91/nhlslackbot/fetch"
)

func main() {
	fmt.Printf("this is a fun project\n")

	if err := fetch.Schedule(); err != nil {
		log.Fatalf("error in fetch schedule: %+v", err)
	}
}
