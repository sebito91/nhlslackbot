package fetch

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

// APIURL is the default API root for NHL stats
const APIURL = "https://statsapi.web.nhl.com/api/v1/"

// Schedule calls out to the NHL API listed at APIURL
// and returns a formatted JSON blob of stats
//
// This function calls the 'schedule' endpoint which
// returns the most recent games by default
// TODO: add options to extend dates
func Schedule() error {
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
