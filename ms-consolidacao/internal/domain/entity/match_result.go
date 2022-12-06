package entity

import "strconv"

type MatchResult struct {
	TeamAScore int
	TeamBScore int
}

func NewMatchResult(teamAScore, teamBScore int) *MatchResult {
	return &MatchResult{
		TeamAScore: teamAScore,
		TeamBScore: teamBScore,
	}
}

func (m *MatchResult) GetResult() string {
	return strconv.Itoa(m.TeamAScore) + "-" + strconv.Itoa(m.TeamBScore)
}
