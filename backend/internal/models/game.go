package models

import (
	"encoding/json"
	"fmt"
)

type Penalty struct {
	PlayerNumber string	`json:"player_number"`
	Time         string	`json:"time"`
}

func (o Penalty) String() string {
	jsonified, _ := json.Marshal(o)
	return fmt.Sprint(string(jsonified))
}

type TeamData struct {
	Score       string			`json:"score"`
	ShotsOnGoal string			`json:"shots_on_goal"`
	Penalties   []Penalty		`json:"penalties"`
}

func (o TeamData) String() string {
	jsonified, _ := json.Marshal(o)
	return fmt.Sprint(string(jsonified))
}

type ScoreboardData struct {
	GameTime string   `json:"game_time"`
	Paused   bool     `json:"paused"`
	Period   string   `json:"period"`
	Home     TeamData `json:"home"`
	Away     TeamData `json:"away"`
}


func (o ScoreboardData) String() string {
	jsonified, _ := json.Marshal(o)
	return fmt.Sprint(string(jsonified))
}


type GameData struct {
	// TODO Eventually we want to add team logos here
	HomeTeam   string         `json:"home_team"`
	AwayTeam   string         `json:"away_team"`
	Scoreboard ScoreboardData `json:"scoreboard"`
}

func (o GameData) String() string {
	jsonified, _ := json.Marshal(o)
	return fmt.Sprint(string(jsonified))
}

