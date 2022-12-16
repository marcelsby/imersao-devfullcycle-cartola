import {
  Autocomplete,
  Avatar,
  Box,
  Button,
  Divider,
  Grid,
  IconButton,
  List,
  ListItem,
  ListItemAvatar,
  ListItemText,
  TextField,
} from "@mui/material"
import { NextPage } from "next"
import Image from "next/image"
import React, { useCallback, useEffect, useMemo, useState } from "react"
import { Label } from "../components/Label"
import { Page } from "../components/Page"
import { Section } from "../components/Section"
import { TeamLogo } from "../components/TeamLogo"
import { Player, PlayersMap } from "../util/models"
import PersonIcon from "@mui/icons-material/Person2"
import DeleteIcon from "@mui/icons-material/Delete"
import { fetcherStats, httpAdmin } from "../util/http"
import { useHttp } from "../hooks/useHttp"

const fakePlayer = {
  id: null,
  name: "Escolha o jogador",
  price: 0,
}

const makeFakePlayer = (key: number) => ({
  ...fakePlayer,
  name: `${fakePlayer.name} ${key + 1}`,
})

const totalPlayers = 4

const fakePlayers: Player[] = new Array(totalPlayers)
  .fill(0)
  .map((_, key) => makeFakePlayer(key))

const ListPlayersPage: NextPage = () => {
  const { data: myPlayers } = useHttp(
    "/my-teams/22087246-01bc-46ad-a9d9-a99a6d734167/players",
    fetcherStats
  )

  const { data: balanceData } = useHttp(
    "/my-teams/22087246-01bc-46ad-a9d9-a99a6d734167/balance",
    fetcherStats,
    { refreshInterval: 5000 }
  )

  const { data: players } = useHttp<Player[]>("/players", fetcherStats, {
    fallbackData: [],
  })

  const [selectedPlayers, setSelectedPlayers] = useState(fakePlayers)

  const countPlayersUsed = useMemo(
    () => selectedPlayers.filter((player) => player.id !== null).length,
    [selectedPlayers]
  )

  const remainingBudget = useMemo(
    () =>
      !balanceData
        ? 0
        : balanceData.balance -
          selectedPlayers.reduce((acc, player) => acc + player.price, 0),
    [selectedPlayers, balanceData]
  )

  useEffect(() => {
    if (!myPlayers) {
      return
    }

    setSelectedPlayers((prev) => {
      console.log(prev.slice(myPlayers.length))
      return [...myPlayers, ...prev.slice(myPlayers.length)]
    })
  }, [myPlayers])

  const addPlayer = useCallback((player: Player) => {
    setSelectedPlayers((prev) => {
      const hasFound = prev.find((p) => p.id === player.id)
      if (hasFound) return prev

      const firstIndexFakePlayer = prev.findIndex((p) => p.id === null)
      if (firstIndexFakePlayer === -1) return prev

      const newPlayers = [...prev]
      newPlayers[firstIndexFakePlayer] = player

      return newPlayers
    })
  }, [])

  const removePlayer = useCallback((index: number) => {
    setSelectedPlayers((prev) => {
      const newPlayers = prev.map((player, key) => {
        if (key === index) {
          return makeFakePlayer(key)
        }
        return player
      })
      return newPlayers
    })
  }, [])

  const saveMyPlayers = useCallback(async () => {
    await httpAdmin.put(
      "/my-teams/22087246-01bc-46ad-a9d9-a99a6d734167/players",
      { players_uuid: selectedPlayers.map((player) => player.id) }
    )
  }, [selectedPlayers])

  return (
    <Page>
      <Grid
        container
        sx={{
          display: "flex",
          justifyContent: "center",
          gap: (theme) => theme.spacing(4),
        }}
      >
        <Grid item xs={12}>
          <Section>
            <TeamLogo
              sx={{
                position: "absolute",
                flexDirection: "row",
                gap: (theme) => theme.spacing(1),
                marginLeft: (theme) => theme.spacing(-1.5),
                marginTop: (theme) => theme.spacing(-3.5),
              }}
            />
            <Box
              sx={{
                display: "flex",
                justifyContent: "flex-end",
                gap: (theme) => theme.spacing(1),
              }}
            >
              <Label>Você ainda tem</Label>
              <Label>C$ {remainingBudget}</Label>
            </Box>
          </Section>
        </Grid>
        <Grid item xs={12}>
          <Section>
            <Grid container>
              <Grid item xs={6}>
                <Autocomplete
                  sx={{ width: 400 }}
                  isOptionEqualToValue={(option, value) => {
                    return option.name
                      .toLowerCase()
                      .includes(value.name.toLowerCase())
                  }}
                  getOptionLabel={(option) => option.name}
                  options={players!}
                  onChange={(_event, newValue) => {
                    if (!newValue) {
                      return
                    }
                    addPlayer(newValue)
                  }}
                  renderOption={(props, option) => (
                    <React.Fragment key={option.name}>
                      <ListItem {...props}>
                        <ListItemAvatar>
                          <Avatar>
                            <Image
                              src={PlayersMap[option.name]}
                              width={40}
                              height={40}
                              alt=""
                            />
                          </Avatar>
                        </ListItemAvatar>
                        <ListItemText
                          primary={`${option.name}`}
                          secondary={`C$ ${option.price}`}
                        />
                      </ListItem>
                      <Divider variant="inset" component="li" />
                    </React.Fragment>
                  )}
                  renderInput={(params) => (
                    <TextField
                      {...params}
                      label="Pesquise um jogador"
                      InputProps={{
                        ...params.InputProps,
                        sx: {
                          backgroundColor: (theme) =>
                            theme.palette.background.default,
                        },
                      }}
                    />
                  )}
                />
              </Grid>
              <Grid item xs={6}>
                <Label>Meu time</Label>
                <List>
                  {selectedPlayers.map((player, key) => (
                    <React.Fragment key={key}>
                      <ListItem
                        secondaryAction={
                          <IconButton
                            edge="end"
                            disabled={player.id === null}
                            onClick={() => removePlayer(key)}
                          >
                            <DeleteIcon />
                          </IconButton>
                        }
                      >
                        <ListItemAvatar>
                          <Avatar>
                            {player.id === null ? (
                              <PersonIcon />
                            ) : (
                              <Image
                                src={PlayersMap[player.name]}
                                width={40}
                                height={40}
                                alt=""
                              />
                            )}
                          </Avatar>
                        </ListItemAvatar>
                        <ListItemText
                          primary={player.name}
                          secondary={`C$ ${player.price}`}
                        />
                      </ListItem>
                      <Divider variant="inset" component="li" />
                    </React.Fragment>
                  ))}
                </List>
              </Grid>
            </Grid>
          </Section>
        </Grid>
        <Grid item>
          <Button
            variant="contained"
            size="large"
            disabled={countPlayersUsed < totalPlayers || remainingBudget < 0}
            onClick={saveMyPlayers}
          >
            Escalar
          </Button>
        </Grid>
      </Grid>
    </Page>
  )
}

export default ListPlayersPage
