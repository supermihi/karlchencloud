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
  } as const,
  main: {
    display: 'flex',
    marginTop: theme.spacing(2),
    flexDirection: 'column',
  } as const,
  buttons: {
    marginBottom: theme.spacing(2),
  } as const,
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
      <Grid container spacing={2} className={classes.buttons}>
        <Grid item xs={4}>
          <Button fullWidth startIcon={<MailOutlineIcon />}>
            Einladung
          </Button>
        </Grid>
        <Grid item xs={4}>
          <Button fullWidth startIcon={<SearchIcon />}>
            Tisch suchen
          </Button>
        </Grid>
        <Grid item xs={4}>
          <Button
            startIcon={<AddIcon />}
            disabled={Boolean(activeTable)}
            fullWidth
            onClick={createTable}
          >
            Neuer Tisch
          </Button>
        </Grid>
      </Grid>
      {activeTable && <TableCard me={me} table={activeTable} startTable={startTable} />}
    </div>
  );
}
