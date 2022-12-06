package service

import (
	"errors"

	"github.com/marcelsby/imersao-devfullcycle-cartola-consolidacao/internal/domain/entity"
)

func ChoosePlayers(myTeam *entity.MyTeam, newChoosenPlayers []entity.Player) error {
	totalCost := 0.0
	totalEarned := 0.0

	for _, player := range newChoosenPlayers {
		if isPlayerInMyTeam(player, myTeam) && !isPlayerInPlayersList(player, &newChoosenPlayers) {
			totalEarned += player.Price
		}

		if !isPlayerInMyTeam(player, myTeam) && isPlayerInPlayersList(player, &newChoosenPlayers) {
			totalCost += player.Price
		}
	}

	if totalCost > myTeam.Score+totalEarned {
		return errors.New("not enough money")
	}

	myTeam.Score += totalEarned - totalCost
	myTeam.Players = []string{}

	for _, player := range newChoosenPlayers {
		myTeam.Players = append(myTeam.Players, player.ID)
	}

	return nil
}

func isPlayerInMyTeam(player entity.Player, myTeam *entity.MyTeam) bool {
	for _, myTeamPlayerID := range myTeam.Players {
		if player.ID == myTeamPlayerID {
			return true
		}
	}

	return false
}

func isPlayerInPlayersList(player entity.Player, playersList *[]entity.Player) bool {
	for _, playersListPlayer := range *playersList {
		if player.ID == playersListPlayer.ID {
			return true
		}
	}

	return false
}
