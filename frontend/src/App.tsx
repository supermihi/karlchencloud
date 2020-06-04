import React from "react";

import "./App.css";
import { Toolbar, makeStyles, Typography, AppBar } from "@material-ui/core";
import LoginPage from "./features/login/LoginPageContainer";
import { selectLogin } from "app/core/login";
import { useSelector } from "react-redux";

const useStyles = makeStyles((theme) => ({
  appBar: {
    position: "relative",
  },
  layout: {
    width: "auto",
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
        {state.loggedIn ? "yes!" : <LoginPage />}
      </main>
    </>
  );
}

export default App;
