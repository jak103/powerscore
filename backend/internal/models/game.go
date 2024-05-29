package models

import "fmt"

type Penalty struct {
	PlayerNumber string
	Time         string
}

func (p Penalty) String() string {
	return fmt.Sprintf("Penalty: {PlayerNumber: %s, Time: %s}", p.PlayerNumber, p.Time)
}

type TeamData struct {
	Score       string
	ShotsOnGoal string
	Penalties   []Penalty
}

func (t TeamData) String() string {
	return fmt.Sprintf("TeamData: {Score: %s, SOG: %s, Penalties: %s}", t.Score, t.ShotsOnGoal, t.Penalties)
}

type ScoreboardData struct {
	GameTime string   `json:"game_time"`
	Paused   bool     `json:"paused"`
	Period   string   `json:"period"`
	Home     TeamData `json:"home"`
	Away     TeamData `json:"away"`
}


func (s ScoreboardData) String() string {
	return fmt.Sprintf("ScoreboardData: {GameTime: %s, Paused: %t, Period: %s, Home: %s, Away: %s}", s.GameTime, s.Paused, s.Period, s.Home, s.Away)
}


type GameData struct {
	// TODO Eventually we want to add team logos here
	HomeTeam   string         `json:"home_team"`
	AwayTeam   string         `json:"away_team"`
	Scoreboard ScoreboardData `json:"scoreboard"`
}

func (g GameData) String() string {
	return fmt.Sprintf("GameData: {HomeTeam: %s, AwayTeam: %s, Scoreboad: %s}", g.HomeTeam, g.AwayTeam, g.Scoreboard)
}

