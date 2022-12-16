import { AppBar, Avatar, Button, Chip, Toolbar } from "@mui/material"
import { Box } from "@mui/system"
import Image from "next/image"
import Link, { LinkProps } from "next/link"
import { useRouter } from "next/router"
import { PropsWithChildren } from "react"
import { useHttp } from "../hooks/useHttp"
import { fetcherStats } from "../util/http"

export type NavbarItemProps = LinkProps & { showUnderline: boolean }

export const NavbarItem = ({
  showUnderline,
  ...linkProps
}: PropsWithChildren<NavbarItemProps>) => {
  return (
    //@ts-expect-error
    <Button
      component={Link}
      sx={{
        color: "white",
        display: "inline-block",
        textAlign: "center",
        "&::after": (theme) => ({
          display: "block",
          content: '""',
          borderBottom: "4px solid",
          borderColor: showUnderline
            ? theme.palette.primary.main
            : "transparent",
          width: "100%",
        }),
      }}
      {...linkProps}
    />
  )
}

export const Navbar = () => {
  const router = useRouter()

  const { data } = useHttp(
    "/my-teams/22087246-01bc-46ad-a9d9-a99a6d734167/balance",
    fetcherStats,
    { refreshInterval: 5000 }
  )

  return (
    <Box sx={{ flexGrow: 1 }}>
      <AppBar
        position="static"
        sx={{
          background: "none",
          boxShadow: "none",
        }}
      >
        <Toolbar>
          <Image
            src="/img/logo.png"
            width={315}
            height={58}
            alt="logo"
            priority={true}
          />
          <Box sx={{ flexGrow: 1, ml: (theme) => theme.spacing(4) }}>
            <NavbarItem href="/" showUnderline={router.pathname === "/"}>
              Home
            </NavbarItem>
            <NavbarItem
              href="/players"
              showUnderline={router.pathname === "/players"}
            >
              Escalação
            </NavbarItem>
            <NavbarItem
              href="/matches"
              showUnderline={router.pathname.startsWith("/matches")}
            >
              Jogo
            </NavbarItem>
          </Box>
          <Chip
            label={data ? data.balance : 0}
            avatar={<Avatar>C$</Avatar>}
            color="secondary"
          />
        </Toolbar>
      </AppBar>
    </Box>
  )
}
