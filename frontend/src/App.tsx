import React from "react";
import { hot } from "react-hot-loader/root";
import { Toolbar, makeStyles, Typography, AppBar } from "@material-ui/core";
import { Component as AuthView } from "./features/auth";
import { useSelector } from "react-redux";
import LobbyView from "features/lobby/Main";
import TableView from "features/table/Main";
import { selectLocation, Location } from "./core/routing";

const useStyles = makeStyles((theme) => ({
  appBar: {
    position: "relative",
  },
  layout: {
    width: "auto",
    backgroundColor: theme.palette.background.paper,
    marginLeft: theme.spacing(2),
    marginRight: theme.spacing(2),
    [theme.breakpoints.up(400 + theme.spacing(2) * 2)]: {
      width: 400,
      marginLeft: "auto",
      marginRight: "auto",
    },
  },
}));

function App() {
  const classes = useStyles();
  const location = useSelector(selectLocation);
  return (
    <>
      <AppBar position="absolute" className={classes.appBar}>
        <Toolbar>
          <Typography variant="h6" color="inherit" noWrap>
            Karlchencloud
          </Typography>
        </Toolbar>
      </AppBar>
      <main className={classes.layout}>
        <Content location={location} />
      </main>
    </>
  );
}
function Content({ location }: { location: Location }) {
  switch (location) {
    case "login":
      return <AuthView />;
    case "lobby":
      return <LobbyView />;
    case "table":
      return <TableView />;
  }
}

export default process.env.NODE_ENV === "development" ? hot(App) : App;
