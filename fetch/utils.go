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
}

// Team is the default struct for a single team
// The League object is made up of multiple 'Team' objects
type Team struct {
}

// Standing is the default struct for a given Team's standing
// The overall league 'Standings' are made up of multiple 'Standing' objects
type Standing struct {
}
