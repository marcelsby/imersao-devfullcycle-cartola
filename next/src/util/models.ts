export type Player = {
  id: string | null
  name: string
  price: number
}

export const PlayersMap: { [key: string]: string } = {
  "Cristiano Ronaldo": "/img/players/Cristiano Ronaldo.png",
  "De Bruyne": "/img/players/De Bruyne.png",
  "Harry Kane": "/img/players/Harry Kane.png",
  Lewandowski: "/img/players/Lewandowski.png",
  Maguirre: "/img/players/Maguirre.png",
  Messi: "/img/players/Messi.png",
  Neymar: "/img/players/Neymar.png",
  Richarlison: "/img/players/Richarlison.png",
  "Vinicius Junior": "/img/players/Vinicius Junior.png",
}

export type Action = {
  playerName: string
  minutes: number
  action: "goal" | "yellow card" | "red card" | "assist"
  score: number
}

export type Match = {
  id: string
  matchDate: string
  teamA: string //Brasil
  teamB: string //Argentina
  score: string //'1-0'
  //score: number;
  actions: Action[]
}

export const TeamsImagesMap: { [key: string]: string } = {
  Alemanha: "/img/flags/Alemanha.png",
  Argentina: "/img/flags/Argentina.png",
  Bélgica: "/img/flags/Belgica.png",
  Brasil: "/img/flags/Brasil.png",
  França: "/img/flags/Franca.png",
  Inglaterra: "/img/flags/Inglaterra.png",
  Polônia: "/img/flags/Polonia.png",
  Portugal: "/img/flags/Portugal.png",
}
