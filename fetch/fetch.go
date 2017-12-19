package fetch

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

// APIURL is the default API root for NHL stats
const APIURL = "https://statsapi.web.nhl.com/api/v1/"

// New returns a new instance of the NHL struct, the default
// type for our score processing
func New() *NHL {
	return &NHL{}
}

// GetSchedule calls out to the NHL API listed at APIURL
// and returns a formatted JSON blob of stats
//
// This function calls the 'schedule' endpoint which
// returns the most recent games by default
// TODO: add options to provide date range
func (n *NHL) GetSchedule() error {
	resp, err := http.Get(fmt.Sprintf("%s/schedule", APIURL))
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	fmt.Printf("DEBUG -- body is: %s\n", body)

	return nil
}

// GetTeams calls the 'teams' endpoint and retrieves the detailed
// information for current teams in the league
//
// TODO: add support for previous seasons
func (n *NHL) GetTeams() error {
	resp, err := http.Get(fmt.Sprintf("%s/teams", APIURL))
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	fmt.Printf("DEBUG -- body is: %s\n", body)

	return nil
}

// GetTeam calls the 'teams' endpoint for a specific ID and retrieves
// the detailed information for current team
//
// TODO: add support for previous seasons
func (n *NHL) GetTeam(id int) error {
	resp, err := http.Get(fmt.Sprintf("%s/teams/%d", APIURL, id))
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	fmt.Printf("DEBUG -- body is: %s\n", body)

	return nil
}
