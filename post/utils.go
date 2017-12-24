package post

import (
	"encoding/json"
	"expvar"
	"fmt"
	"net/http"
	"net/http/pprof"
	"strings"
	"sync"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/cors"
	"github.com/nlopes/slack"
	"github.com/sebito91/nhlslackbot/fetch"
)

// PostMap is a global map to handle callbacks depending on the provided user
// This mapping stores off the userID to reply to
var PostMap map[string]string

// PostLock is the complement for the global PostMap to ensure concurrent
// access doesn't race
var PostLock sync.RWMutex

// IndexHandler returns data to the default port
func indexHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte(fmt.Sprintf("incorrect path: %s", r.URL.Path)))
		return
	}

	switch r.Method {
	case "GET":
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(fmt.Sprintf("%v", `¯\_(ツ)_/¯ GET`)))
		return
	case "POST":
		w.WriteHeader(http.StatusMovedPermanently)
		w.Write([]byte("cannot post to this endpoint"))
		return
	default:
	}
}

func postHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte(fmt.Sprintf("incorrect path: %s", r.URL.Path)))
		return
	}

	if r.Body == nil {
		w.WriteHeader(http.StatusNotAcceptable)
		w.Write([]byte("empty body"))
		return
	}
	defer r.Body.Close()

	err := r.ParseForm()
	if err != nil {
		w.WriteHeader(http.StatusGone)
		w.Write([]byte("could not parse body"))
		return
	}

	// slack API calls the data POST a 'payload'
	reply := r.PostFormValue("payload")
	if len(reply) == 0 {
		w.WriteHeader(http.StatusNoContent)
		w.Write([]byte("could not find payload"))
		return
	}

	var payload slack.AttachmentActionCallback
	err = json.NewDecoder(strings.NewReader(reply)).Decode(&payload)
	if err != nil {
		w.WriteHeader(http.StatusGone)
		w.Write([]byte("could not process payload"))
		return
	}

	action := payload.Actions[0].Value
	switch action {
	case "yes":
		grabStats(w, r)
	case "no":
		w.Write([]byte("No worries, let me know later on if you do!"))
	default:
		w.WriteHeader(http.StatusNotAcceptable)
		w.Write([]byte(fmt.Sprintf("could not process callback: %s", action)))
		return
	}

	w.WriteHeader(http.StatusOK)
}

// askIntent is the initial request back to user if they'd like to see
// the scores from the most recent slate of games
//
// NOTE: This is a contrived example of the functionality, but ideally here
// we would ask users to specify a date, or maybe a team, or even
// a specific game which we could present back
func (s *Slack) askIntent(ev *slack.MessageEvent) error {
	params := slack.NewPostEphemeralParameters()
	attachment := slack.Attachment{
		Text:       "Would you like to see the most recent scores?",
		CallbackID: fmt.Sprintf("ask_%s", ev.User),
		Color:      "#666666",
		Actions: []slack.AttachmentAction{
			slack.AttachmentAction{
				Name:  "action",
				Text:  "No thanks!",
				Type:  "button",
				Value: "no",
			},
			slack.AttachmentAction{
				Name:  "action",
				Text:  "Yes, please!",
				Type:  "button",
				Value: "yes",
			},
		},
	}

	params.Attachments = []slack.Attachment{attachment}
	params.User = ev.User
	params.AsUser = true

	_, err := s.Client.PostEphemeral(
		ev.Channel,
		ev.User,
		slack.MsgOptionAttachments(params.Attachments...),
		slack.MsgOptionPostEphemeralParameters(params),
	)
	if err != nil {
		return err
	}

	return nil
}

// NewHandler instantiaties the web handler for listening on the API
func (s *Slack) NewHandler() (http.Handler, error) {
	r := chi.NewRouter()

	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)

	r.Use(middleware.NoCache)
	r.Use(middleware.Heartbeat("/ping"))

	cors := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		AllowCredentials: true,
		MaxAge:           300, // Maximum value not ignored by any of major browsers
	})
	r.Use(cors.Handler)

	r.Get("/", indexHandler)
	r.Post("/", postHandler)

	r.Get("/debug/pprof/*", pprof.Index)
	r.Get("/debug/vars", func(w http.ResponseWriter, r *http.Request) {
		first := true
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		fmt.Fprintf(w, "{\n")
		expvar.Do(func(kv expvar.KeyValue) {
			if !first {
				fmt.Fprintf(w, ",\n")
			}
			first = false
			fmt.Fprintf(w, "%q: %s", kv.Key, kv.Value)
		})
		fmt.Fprintf(w, "\n}\n")
	})

	return r, nil
}

// grabStats will process the information from the API and return the data to
// our user!
func grabStats(w http.ResponseWriter, r *http.Request) {
	n := fetch.New()

	buf, err := n.GetSchedule()
	if err != nil {
		w.WriteHeader(http.StatusNoContent)
		w.Write([]byte(fmt.Sprintf("error processing schedule; %v", err)))
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(buf)
}
