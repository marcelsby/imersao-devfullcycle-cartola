package entity

import "github.com/google/uuid"

type GameAction struct {
	ID         string
	PlayerID   string
	PlayerName string
	TeamID     string
	Minute     int
	Action     string
	Score      int
}

func NewGameAction(playerID string, teamID string, minute int, action string, score int) *GameAction {
	return &GameAction{
		ID:       uuid.New().String(),
		PlayerID: playerID,
		TeamID:   teamID,
		Minute:   minute,
		Action:   action,
		Score:    score,
	}
}
