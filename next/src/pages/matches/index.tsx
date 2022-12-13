import { Box } from "@mui/material"
import { NextPage } from "next"
import { MatchResult } from "../../components/MatchResult"
import { Page } from "../../components/Page"

const ListMatchesPage: NextPage = () => {
  return (
    <Page>
      <Box
        sx={{
          display: "flex",
          flexDirection: "column",
          alignItems: "center",
          gap: (theme) => theme.spacing(3),
        }}
      >
        <MatchResult match={{ teamA: "Brasil", teamB: "Argentina" }} />
        <MatchResult match={{ teamA: "Brasil", teamB: "Argentina" }} />
        <MatchResult match={{ teamA: "Brasil", teamB: "Argentina" }} />
      </Box>
    </Page>
  )
}

export default ListMatchesPage
