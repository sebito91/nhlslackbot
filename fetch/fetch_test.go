// +build small

package fetch

var payload = `
{
	copyright: "NHL and the NHL Shield are registered trademarks of the National Hockey League. NHL and NHL team marks are the property of the NHL and its teams. © NHL 2017. All Rights Reserved.",
	totalItems: 5,
	totalEvents: 0,
	totalGames: 5,
	totalMatches: 0,
	wait: 10,
	dates: [
	{
		date: "2017-12-18",
		totalItems: 5,
		totalEvents: 0,
		totalGames: 5,
		totalMatches: 0,
		games: [
		{
			gamePk: 2017020511,
			link: "/api/v1/game/2017020511/feed/live",
			gameType: "R",
			season: "20172018",
			gameDate: "2017-12-19T00:00:00Z",
			status: {
				abstractGameState: "Final",
				codedGameState: "7",
				detailedState: "Final",
				statusCode: "7",
				startTimeTBD: false
			},
			teams: {
				away: {
					leagueRecord: {
						wins: 20,
						losses: 13,
						ot: 1,
						type: "league"
					},
					score: 2,
					team: {
						id: 29,
						name: "Columbus Blue Jackets",
						link: "/api/v1/teams/29"
					}
				},
				home: {
					leagueRecord: {
						wins: 16,
						losses: 10,
						ot: 5,
						type: "league"
					},
					score: 7,
					team: {
						id: 6,
						name: "Boston Bruins",
						link: "/api/v1/teams/6"
					}
				}
			},
			venue: {
				name: "TD Garden",
				link: "/api/v1/venues/null"
			},
			content: {
				link: "/api/v1/game/2017020511/content"
			}
		},
		{
			gamePk: 2017020512,
			link: "/api/v1/game/2017020512/feed/live",
			gameType: "R",
			season: "20172018",
			gameDate: "2017-12-19T00:00:00Z",
			status: {
				abstractGameState: "Final",
				codedGameState: "7",
				detailedState: "Final",
				statusCode: "7",
				startTimeTBD: false
			},
			teams: {
				away: {
					leagueRecord: {
						wins: 14,
						losses: 12,
						ot: 8,
						type: "league"
					},
					score: 3,
					team: {
						id: 24,
						name: "Anaheim Ducks",
						link: "/api/v1/teams/24"
					}
				},
				home: {
					leagueRecord: {
						wins: 19,
						losses: 9,
						ot: 5,
						type: "league"
					},
					score: 5,
					team: {
						id: 1,
						name: "New Jersey Devils",
						link: "/api/v1/teams/1"
					}
				}
			},
			venue: {
				name: "Prudential Center",
				link: "/api/v1/venues/null"
			},
			content: {
				link: "/api/v1/game/2017020512/content"
			}
		},
		{
			gamePk: 2017020513,
			link: "/api/v1/game/2017020513/feed/live",
			gameType: "R",
			season: "20172018",
			gameDate: "2017-12-19T00:00:00Z",
			status: {
				abstractGameState: "Final",
				codedGameState: "7",
				detailedState: "Final",
				statusCode: "7",
				startTimeTBD: false
			},
			teams: {
				away: {
					leagueRecord: {
						wins: 21,
						losses: 10,
						ot: 4,
						type: "league"
					},
					score: 4,
					team: {
						id: 26,
						name: "Los Angeles Kings",
						link: "/api/v1/teams/26"
					}
				},
				home: {
					leagueRecord: {
						wins: 14,
						losses: 12,
						ot: 7,
						type: "league"
					},
					score: 1,
					team: {
						id: 4,
						name: "Philadelphia Flyers",
						link: "/api/v1/teams/4"
					}
				}
			},
			venue: {
				name: "Wells Fargo Center",
				link: "/api/v1/venues/null"
			},
			content: {
				link: "/api/v1/game/2017020513/content"
			}
		},
		{
			gamePk: 2017020514,
			link: "/api/v1/game/2017020514/feed/live",
			gameType: "R",
			season: "20172018",
			gameDate: "2017-12-19T02:00:00Z",
			status: {
				abstractGameState: "Final",
				codedGameState: "7",
				detailedState: "Final",
				statusCode: "7",
				startTimeTBD: false
			},
			teams: {
				away: {
					leagueRecord: {
						wins: 17,
						losses: 15,
						ot: 3,
						type: "league"
					},
					score: 2,
					team: {
						id: 5,
						name: "Pittsburgh Penguins",
						link: "/api/v1/teams/5"
					}
				},
				home: {
					leagueRecord: {
						wins: 16,
						losses: 15,
						ot: 2,
						type: "league"
					},
					score: 4,
					team: {
						id: 21,
						name: "Colorado Avalanche",
						link: "/api/v1/teams/21"
					}
				}
			},
			venue: {
				name: "Pepsi Center",
				link: "/api/v1/venues/null"
			},
			content: {
				link: "/api/v1/game/2017020514/content"
			}
		},
		{
			gamePk: 2017020515,
			link: "/api/v1/game/2017020515/feed/live",
			gameType: "R",
			season: "20172018",
			gameDate: "2017-12-19T02:00:00Z",
			status: {
				abstractGameState: "Final",
				codedGameState: "7",
				detailedState: "Final",
				statusCode: "7",
				startTimeTBD: false
			},
			teams: {
				away: {
					leagueRecord: {
						wins: 17,
						losses: 11,
						ot: 4,
						type: "league"
					},
					score: 3,
					team: {
						id: 28,
						name: "San Jose Sharks",
						link: "/api/v1/teams/28"
					}
				},
				home: {
					leagueRecord: {
						wins: 15,
						losses: 17,
						ot: 2,
						type: "league"
					},
					score: 5,
					team: {
						id: 22,
						name: "Edmonton Oilers",
						link: "/api/v1/teams/22"
					}
				}
			},
			venue: {
				name: "Rogers Place",
				link: "/api/v1/venues/null"
			},
			content: {
				link: "/api/v1/game/2017020515/content"
			}
		}
		],
		events: [ ],
		matches: [ ]
	}
	]
}
`
