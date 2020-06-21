import React from 'react';
import { Toolbar, makeStyles, Typography, AppBar } from '@material-ui/core';
import Login from 'features/auth/Login';
import Register from 'features/auth/Register';
import { connect } from 'react-redux';
import Lobby from 'features/lobby/Lobby';
import TableView from 'features/table/Main';
import { selectLocation, Location } from './app/routing';
import { createSelector } from '@reduxjs/toolkit';

const useStyles = makeStyles((theme) => ({
  appBar: {
    position: 'relative',
  },
  layout: {
    width: 'auto',
    backgroundColor: theme.palette.background.paper,
    marginLeft: theme.spacing(2),
    marginRight: theme.spacing(2),
    [theme.breakpoints.up(400 + theme.spacing(2) * 2)]: {
      width: 400,
      marginLeft: 'auto',
      marginRight: 'auto',
    },
  },
}));

interface Props {
  location: Location;
}
function AppView({ location }: Props) {
  const classes = useStyles();
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
    case Location.login:
      return <Login />;
    case Location.register:
      return <Register />;
    case Location.lobby:
      return <Lobby />;
    case Location.table:
      return <TableView />;
  }
}
const mapState = createSelector(selectLocation, (location) => ({ location }));
export default connect(mapState)(AppView);
