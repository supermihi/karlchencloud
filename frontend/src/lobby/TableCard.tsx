import * as React from 'react';

import Card from '@material-ui/core/Card';
import CardContent from '@material-ui/core/CardContent';
import CardActions from '@material-ui/core/CardActions';
import Button from '@material-ui/core/Button';
import List from '@material-ui/core/List';

import makeStyles from '@material-ui/core/styles/makeStyles';
import Typography from '@material-ui/core/Typography';

import { canStartTable, canContinueTable, waitingForPlayers, Table } from 'model/table';
import { TablePhase } from 'api/karlchen_pb';
import InviteDialog from './InviteLinkDialog';
import PlayerItem from './PlayerItem';
import { User } from 'model/core';

interface Props {
  table: Table;
  me: User;
  startTable: () => void;
}

const useStyles = makeStyles((theme) => ({
  continueButton: {
    marginLeft: 'auto',
  } as const,
  note: {
    color: theme.palette.text.hint,
  } as const,
}));

export default function CurrentTableView({ table, me, startTable }: Props): React.ReactElement {
  const classes = useStyles();
  const [inviteOpen, setInviteOpen] = React.useState(false);

  const started = table.phase !== TablePhase.NOT_STARTED;
  const owner = table.members.find((p) => p.id === table.owner);

  return (
    <Card variant="outlined">
      <CardContent>
        <Typography variant="h4" component="h2">{`${owner?.name}'s Tisch`}</Typography>
        <Typography color="textSecondary">Erstellt {table.created.toLocaleString()}</Typography>
        <List>
          {table.members.map((player) => (
            <PlayerItem me={player.id === me.id} player={player} key={player.id} />
          ))}
        </List>
        {waitingForPlayers(table) && (
          <Typography variant="body1" className={classes.note}>
            Warte auf Mitspieler …
          </Typography>
        )}
      </CardContent>
      <CardActions>
        {!started && (
          <>
            <Button onClick={() => setInviteOpen(true)}>Mitspieler einladen</Button>
            <InviteDialog
              open={inviteOpen}
              handleClose={() => setInviteOpen(false)}
              inviteCode={table.invite || ''}
            />
          </>
        )}
        <div style={{ flexGrow: 1 }} />
        {canContinueTable(table) && (
          <Button variant="contained" color="primary" size="small">
            Weiter spielen
          </Button>
        )}
        {!started && (
          <Button
            variant="contained"
            color="primary"
            disabled={!canStartTable(table, me)}
            onClick={startTable}
          >
            Starten
          </Button>
        )}
      </CardActions>
    </Card>
  );
}
