import React from "react";
import { hot } from "react-hot-loader/root";
import { Toolbar, makeStyles, Typography, AppBar } from "@material-ui/core";
import Welcome from "./features/login";
import { selectLogin } from "core/login";
import { useSelector } from "react-redux";
import RoomView from "features/room/RoomContainer";

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
  const state = useSelector(selectLogin);
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
        {state.loggedIn ? <RoomView /> : <Welcome />}
      </main>
    </>
  );
}

export default process.env.NODE_ENV === "development" ? hot(App) : App;
