import makeStyles from '@material-ui/core/styles/makeStyles';
import React from 'react';
import { useSelector } from 'react-redux';
import { selectLocation, Location } from 'routing';
import LoginWithToken from 'session/view/LoginWithToken';
import Lobby from 'lobby/Lobby';
import Play from 'play/view/Play';
import LoginOrRegister from './session/view/LoginOrRegister';

export default function App(): React.ReactElement {
  const location = useSelector(selectLocation);
  const classes = useStyles();
  return (
    <>
      <main className={classes.layout}>
        <Content location={location} />
      </main>
    </>
  );
}
function Content({ location }: { location: Location }) {
  switch (location) {
    case Location.login:
      return <LoginOrRegister />;
    case Location.loginWithToken:
      return <LoginWithToken />;
    case Location.lobby:
      return <Lobby />;
    case Location.table:
      return <Play />;
  }
}

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
