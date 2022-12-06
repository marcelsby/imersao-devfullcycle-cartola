package entity

import "time"

type Match struct {
	ID      string
	TeamA   *Team
	TeamB   *Team
	TeamAID string
	TeamBID string
	Date    time.Time
	Status  string
	Result  MatchResult
	Actions []GameAction
}

func NewMatch(id string, teamA *Team, teamB *Team, date time.Time) *Match {
	return &Match{
		ID:      id,
		TeamA:   teamA,
		TeamB:   teamB,
		TeamAID: teamA.ID,
		TeamBID: teamB.ID,
		Date:    date,
	}
}
