import React from 'react';
import { Table } from 'model/table';
import AddIcon from '@material-ui/icons/Add';
import SearchIcon from '@material-ui/icons/Search';
import MailOutlineIcon from '@material-ui/icons/MailOutline';
import Grid from '@material-ui/core/Grid';
import Button from '@material-ui/core/Button';
import makeStyles from '@material-ui/core/styles/makeStyles';
import CurrentTableView from './CurrentTableView';
import AcceptInviteDialog from './AcceptInviteDialog';
import { User } from 'model/core';

interface Props {
  activeTable: Table | null;
  me: User;
  createTable: () => void;
  suppliedInviteCode: string | null;
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

export default ({ activeTable, createTable, me, suppliedInviteCode }: Props) => {
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
      {suppliedInviteCode && <AcceptInviteDialog />}
      {activeTable && <CurrentTableView me={me} table={activeTable} />}
    </div>
  );
};
