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
    flexGrow: 1,
    backgroundColor: theme.palette.background.paper,
    marginLeft: theme.spacing(2),
    marginRight: theme.spacing(2),
    [theme.breakpoints.up(800 + theme.spacing(2) * 2)]: {
      width: 800,
      marginLeft: 'auto',
      marginRight: 'auto',
    },
  } as const,
  '@global': {
    body: {
      margin: 0,
    },
    '#root': {
      height: '100vh',
      width: '100vw',
      display: 'flex',
      flexDirection: 'column',
    },
  },
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
