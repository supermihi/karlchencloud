import React from 'react';
import { Table } from 'model/table';
import AddIcon from '@material-ui/icons/Add';
import SearchIcon from '@material-ui/icons/Search';
import MailOutlineIcon from '@material-ui/icons/MailOutline';
import Grid from '@material-ui/core/Grid';
import Button from '@material-ui/core/Button';
import makeStyles from '@material-ui/core/styles/makeStyles';
import TableCard from './TableCard';
import { User } from 'model/core';
import AppBar from '../shared/AppBar';
import { BottomNavigation, BottomNavigationAction, Paper } from '@material-ui/core';

interface Props {
  activeTable: Table | null;
  me: User;
  createTable: () => void;
  startTable: () => void;
}

const useStyles = makeStyles((theme) => ({
  addTable: {
    marginTop: theme.spacing(2),
    alignSelf: 'center',
  },
  main: {
    display: 'flex',
    flexDirection: 'column',
  },
  content: {
    marginTop: theme.spacing(2),
    marginBottom: theme.spacing(2),
    flexGrow: 1,
  },
  buttons: {
    marginBottom: theme.spacing(2),
  },
}));

export default function LobbyView({
  activeTable,
  createTable,
  me,
  startTable,
}: Props): React.ReactElement {
  const classes = useStyles();
  return (
    <div className={classes.main}>
      <AppBar location="Lobby" />
      <Paper className={classes.content}>
        <div>
          {activeTable && <TableCard me={me} table={activeTable} startTable={startTable} />}
        </div>
      </Paper>
      <BottomNavigation showLabels>
        <BottomNavigationAction
          label="Neu"
          disabled={Boolean(activeTable)}
          onClick={createTable}
          icon={<AddIcon />}
        />
        <BottomNavigationAction label="Suchen" icon={<SearchIcon />} />
        <BottomNavigationAction label="Einladung" icon={<MailOutlineIcon />} />
      </BottomNavigation>
    </div>
  );
}
