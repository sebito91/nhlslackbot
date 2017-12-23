package post

import (
	"context"
	"fmt"
	"os"

	"github.com/nlopes/slack"
	"go.uber.org/zap"
)

// Slack is the primary struct for our slackbot
type Slack struct {
	Name  string
	Token string

	User   string
	UserID string

	Logger *zap.SugaredLogger

	Client *slack.Client
}

// New returns a new instance of the Slack struct, primary for our slackbot
func New() (*Slack, error) {
	token := os.Getenv("SLACK_TOKEN")
	if len(token) == 0 {
		return nil, fmt.Errorf("could not discover API token")
	}

	return &Slack{Client: slack.New(token), Token: token, Name: "nhlslackbot"}, nil
}

// Run is the primary service to generate and kick off the slackbot listener
// This portion receives all incoming Real Time Messages notices from the workspace
// as registered by the API token
func (s *Slack) Run(ctx context.Context) error {
	authTest, err := s.Client.AuthTest()
	if err != nil {
		return fmt.Errorf("did not authenticate: %+v", err)
	}

	s.User = authTest.User
	s.UserID = authTest.UserID

	go s.run(ctx)
	return nil
}

func (s *Slack) run(ctx context.Context) {
	s.Client.SetDebug(true)

	rtm := s.Client.NewRTM()
	go rtm.ManageConnection()

	for msg := range rtm.IncomingEvents {
		s.Logger.Info("Event Received: ")

		switch ev := msg.Data.(type) {
		case *slack.MessageEvent:
			s.Logger.Infof("Message: %v\n", ev)
		case *slack.PresenceChangeEvent:
			s.Logger.Infof("Presence Change: %v\n", ev)
		case *slack.RTMError:
			s.Logger.Infof("Error: %s\n", ev.Error())
		default:
			// Ignore other events..
			s.Logger.Infof("Unexpected: %v\n", msg.Data)
		}
	}
}
