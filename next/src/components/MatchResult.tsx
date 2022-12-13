import { styled, Typography } from "@mui/material"
import Box from "@mui/material/Box"
import Image from "next/image"
import { Match, TeamsImagesMap } from "../util/models"

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
      <Flag src={TeamsImagesMap[match.teamA]} alt={match.teamA} />
      <ResultContainer>
        <ResultItem width="150px" justifyContent="flex-end">
          <Typography variant="h6">Brasil</Typography>
        </ResultItem>
        <ResultItem width="100px" justifyContent="center" position="relative">
          <Box sx={{ position: "absolute", bottom: 0.5, fontSize: ".7rem" }}>
            12/12/2022 00:00
          </Box>
          <Typography variant="h6" sx={{ fontWeight: "900" }}>
            1-0
          </Typography>
        </ResultItem>
        <ResultItem width="150px" justifyContent="flex-start">
          <Typography variant="h6">Argentina</Typography>
        </ResultItem>
      </ResultContainer>
      <Flag src={TeamsImagesMap[match.teamB]} alt={match.teamB} />
    </Box>
  )
}
