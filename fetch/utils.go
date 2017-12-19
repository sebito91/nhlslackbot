package fetch

// NHL is the default struct for the data parsed from APIURL
type NHL struct {
	Copyright string `json:"copyright"`

	Schedule  ScheduleDate
	League    []Team
	Standings []Standing
}

// ScheduleDate lists out the individual games in a given date range
type ScheduleDate struct {
	TotalItems int `json:"totalItems"`
	TotalGames int `json:"totalGames"`
	Wait       int `json:"wait"`

	Dates []Date `json:"dates"`
}

// Date is the holder struct for an individual date in the query
// for historic data
type Date struct {
	Date       string `json:"date"`
	TotalItems int    `json:"totalItems"`
	TotalGames int    `json:"totalGames"`

	Games []Game `json:"games"`
}

// Game is the default struct for a single game
// A Schedule object is typically made up of multiple 'Game' objects
type Game struct {
	ID     int    `json:"gamePk"`
	Link   string `json:"link"`
	Type   string `json:"gameType"`
	Season string `json:"season"`
	Date   string `json:"gameDate"`

	Status struct {
		AbtractState  string `json:"abstractGameState"`
		CodedState    string `json:"codedGameState"`
		DetailedState string `json:"detailedState"`
		StatusCode    string `json:"statsCode"`
		StartTimeTBD  bool   `json:"startTimeTBD"`
	} `json:"status"`

	Teams struct {
		Away HomeAway `json:"away"`
		Home HomeAway `json:"home"`
	} `json:"teams"`

	Venue struct {
		Name string `json:"name"`
		Link string `json:"link"`
	} `json:"venue"`

	Content struct {
		Link string `json:"link"`
	} `json:"content"`
}

// HomeAway is an inner struct for the scores of a given game
type HomeAway struct {
	LeagueRecord struct {
		Wins   int    `json:"wins"`
		Losses int    `json:"losses"`
		OT     int    `json:"ot"`
		Type   string `json:"league"`
	} `json:"leagueRecord"`

	Score int `json:"score"`

	Team struct {
		ID   int    `json:"id"`
		Name string `json:"name"`
		Link string `json:"link"`
	} `json:"team"`
}

// Team is the default struct for a single team
// The League object is made up of multiple 'Team' objects
type Team struct {
}

// Standing is the default struct for a given Team's standing
// The overall league 'Standings' are made up of multiple 'Standing' objects
type Standing struct {
}
