import * as React from 'react';
import makeStyles from '@material-ui/core/styles/makeStyles';
import Login from 'features/auth/Login';
import Register from 'features/auth/Register';
import Lobby from 'features/lobby/Lobby';
import Table from 'features/table/Main';
import { Location } from '../app/routing';
import AppBar from './AppBar';
const useStyles = makeStyles((theme) => ({
  layout: {
    width: 'auto',
    backgroundColor: theme.palette.background.paper,
    marginLeft: theme.spacing(2),
    marginRight: theme.spacing(2),
    [theme.breakpoints.up(600 + theme.spacing(2) * 2)]: {
      width: 600,
      marginLeft: 'auto',
      marginRight: 'auto',
    },
  } as const,
}));

interface Props {
  location: Location;
}

export default function AppView({ location }: Props) {
  const classes = useStyles();
  return (
    <>
      <AppBar />
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
      return <Table />;
  }
}
