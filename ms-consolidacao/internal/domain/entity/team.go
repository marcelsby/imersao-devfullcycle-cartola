package entity

type Team struct {
	ID      string
	Name    string
	Players []*Player
}

func NewTeam(id, name string) *Team {
	return &Team{
		ID:   id,
		Name: name,
	}
}

func (team *Team) AddPlayer(player *Player) {
	team.Players = append(team.Players, player)
}

func (team *Team) RemovePlayer(player *Player) {
	for i, p := range team.Players {
		if p.ID == player.ID {
			team.Players = append(team.Players[:i], team.Players[i+1:]...)
			return
		}
	}
}
