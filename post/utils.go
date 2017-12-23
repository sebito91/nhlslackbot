package post

import (
	"fmt"
	"net/http"

	"github.com/nlopes/slack"
)

// IndexHandler returns data to the default port
func IndexHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte(fmt.Sprintf("incorrect path: %s", r.URL.Path)))
		return
	}

	switch r.Method {
	case "GET":
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(fmt.Sprintf("%v", `¯\_(ツ)_/¯`)))
		return
	case "POST":
		postHandler(w, r)
	default:
	}
}

func postHandler(w http.ResponseWriter, r *http.Request) {

}

// askIntent is the initial request back to user if they'd like to see
// the scores from the most recent slate of games
//
// NOTE: This is a contrived example of the functionality, but ideally here
// we would ask users to specify a date, or maybe a team, or even
// a specific game which we could present back
func (s *Slack) askIntent(ev *slack.MessageEvent) error {
	s.Logger.Printf("[DEBUG] would print out fun here")
	return nil
}
