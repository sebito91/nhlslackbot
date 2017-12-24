package post

import (
	"context"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/nlopes/slack"
)

// Slack is the primary struct for our slackbot
type Slack struct {
	Name  string
	Token string

	User   string
	UserID string

	Logger *log.Logger

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

	s.Logger.Printf("[INFO]  bot is now registered as %s (%s)\n", s.User, s.UserID)

	go s.run(ctx)
	return nil
}

func (s *Slack) run(ctx context.Context) {
	slack.SetLogger(s.Logger)
	//	s.Client.SetDebug(true)

	rtm := s.Client.NewRTM()
	go rtm.ManageConnection()

	s.Logger.Printf("[INFO]  now listening for incoming messages...")
	for msg := range rtm.IncomingEvents {
		switch ev := msg.Data.(type) {
		case *slack.MessageEvent:
			if len(ev.User) == 0 {
				continue
			}

			// check if we have a DM, or standard channel post
			direct := strings.HasPrefix(ev.Msg.Channel, "D")

			if !direct && !strings.Contains(ev.Msg.Text, "@"+s.UserID) {
				// msg not for us!
				continue
			}

			user, err := s.Client.GetUserInfo(ev.User)
			if err != nil {
				s.Logger.Printf("[WARN]  could not grab user information: %s", ev.User)
				continue
			}

			s.Logger.Printf("[DEBUG] received message from %s (%s)\n", user.Profile.RealName, ev.User)

			err = s.askIntent(ev)
			if err != nil {
				s.Logger.Printf("[ERROR] posting ephemeral reply to user (%s): %+v\n", ev.User, err)
			}
		case *slack.RTMError:
			s.Logger.Printf("[ERROR] %s\n", ev.Error())
		}
	}
}
