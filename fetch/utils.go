package fetch

// NHL is the default struct for the data parsed from APIURL
type NHL struct {
	Schedule  []Game
	League    []Team
	Standings []Standing
}

// Game is the default struct for a single game
// A Schedule object is typically made up of multiple 'Game' objects
type Game struct {
	ID     string `json:"gamePk"`
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
}

// HomeAway is an inner struct for the scores of a given game
type HomeAway struct {
	LeagueRecord struct {
		Wins   string `json:"wins"`
		Losses string `json:"losses"`
		OT     string `json:"ot"`
		Type   string `json:"league"`
	} `json:"leagueRecord"`

	Score int `json:"score"`

	Team struct {
		ID   int    `json:"id"`
		Name string `json:"name"`
		Link string `json:"link"`
	} `json:"team"`

	Venue struct {
		Name string `json:"name"`
		Link string `json:"link"`
	} `json:"venue"`

	Content struct {
		Link string `json:"link"`
	} `json:"content"`
}

// Team is the default struct for a single team
// The League object is made up of multiple 'Team' objects
type Team struct {
}

// Standing is the default struct for a given Team's standing
// The overall league 'Standings' are made up of multiple 'Standing' objects
type Standing struct {
}

//		{
//			gamePk: 2017020511,
//			link: "/api/v1/game/2017020511/feed/live",
//			gameType: "R",
//			season: "20172018",
//			gameDate: "2017-12-19T00:00:00Z",
//			status: {
//				abstractGameState: "Final",
//				codedGameState: "7",
//				detailedState: "Final",
//				statusCode: "7",
//				startTimeTBD: false
//			},
//			teams: {
//				away: {
//					leagueRecord: {
//						wins: 20,
//						losses: 13,
//						ot: 1,
//						type: "league"
//					},
//					score: 2,
//					team: {
//						id: 29,
//						name: "Columbus Blue Jackets",
//						link: "/api/v1/teams/29"
//					}
//				},
//				home: {
//					leagueRecord: {
//						wins: 16,
//						losses: 10,
//						ot: 5,
//						type: "league"
//					},
//					score: 7,
//					team: {
//						id: 6,
//						name: "Boston Bruins",
//						link: "/api/v1/teams/6"
//					}
//				}
//			},
//			venue: {
//				name: "TD Garden",
//				link: "/api/v1/venues/null"
//			},
//			content: {
//				link: "/api/v1/game/2017020511/content"
//			}
//		},
