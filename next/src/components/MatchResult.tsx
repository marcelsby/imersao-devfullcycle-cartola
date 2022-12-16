import { styled, Typography } from "@mui/material"
import Box from "@mui/material/Box"
import Image from "next/image"
import { Match, TeamsImagesMap } from "../util/models"
import { format, parseISO } from "date-fns"

const ResultContainer = styled(Box)(({ theme }) => ({
  display: "flex",
  width: "400px",
  backgroundColor: theme.palette.background.default,
  alignItem: "center",
  padding: 0,
  border: "none !important",
  boxShadow: "none",
}))

const ResultItem = styled(Box)(() => ({
  height: "75px",
  display: "flex",
  alignItems: "center",
}))

type FlagProps = {
  src: string
  alt: string
}

const Flag = ({ src, alt }: FlagProps) => {
  return (
    <Image
      src={src}
      alt={alt}
      width={121}
      height={76}
      style={{
        marginLeft: "-5px",
        marginRight: "-5px",
      }}
    />
  )
}

type MatchResultProps = {
  match: Match
}

export const MatchResult = ({ match }: MatchResultProps) => {
  return (
    <Box display="flex">
      <Flag src={TeamsImagesMap[match.team_a]} alt={match.team_a} />
      <ResultContainer>
        <ResultItem width="150px" justifyContent="flex-end">
          <Typography variant="h6">{match.team_a}</Typography>
        </ResultItem>
        <ResultItem width="100px" justifyContent="center" position="relative">
          <Box sx={{ position: "absolute", bottom: 0.5, fontSize: ".7rem" }}>
            {format(parseISO(match.match_date), "dd/MM/yyyy HH:mm")}
          </Box>
          <Typography variant="h6" sx={{ fontWeight: "900" }}>
            {match.result}
          </Typography>
        </ResultItem>
        <ResultItem width="150px" justifyContent="flex-start">
          <Typography variant="h6">{match.team_b}</Typography>
        </ResultItem>
      </ResultContainer>
      <Flag src={TeamsImagesMap[match.team_b]} alt={match.team_b} />
    </Box>
  )
}
