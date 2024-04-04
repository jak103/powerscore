package models

type Penalty struct {
	PlayerNumber string
	Time         string
}

type TeamData struct {
	Score       string
	ShotsOnGoal string
	Penalties   []Penalty
}

type ScoreboardData struct {
	GameTime string   `json:"game_time"`
	Paused   bool     `json:"paused"`
	Period   string   `json:"period"`
	Home     TeamData `json:"home"`
	Away     TeamData `json:"away"`
}

type GameData struct {
	// TODO Eventually we want to add team logos here
	HomeTeam   string         `json:"home_team"`
	AwayTeam   string         `json:"away_team"`
	Scoreboard ScoreboardData `json:"scoreboard"`
}
