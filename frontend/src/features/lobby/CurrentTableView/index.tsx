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
import GrowDiv from 'components/GrowDiv';
import InviteDialog from '../InviteLinkDialog';
import PlayerItem from './PlayerItem';
import { User } from 'model/core';
import { useDispatch } from 'react-redux';
import { startTable } from 'app/game/table';

interface Props {
  table: Table;
  me: User;
}
const useStyles = makeStyles((theme) => ({
  continueButton: {
    marginLeft: 'auto',
  } as const,
  note: {
    color: theme.palette.text.hint,
  } as const,
}));

export default function CurrentTableView({ table, me }: Props) {
  const dispatch = useDispatch();
  const owner = table.players.find((p) => p.id === table.owner);
  const classes = useStyles();
  const [inviteOpen, setInviteOpen] = React.useState(false);
  const started = table.phase !== TablePhase.NOT_STARTED;
  return (
    <Card variant="outlined">
      <CardContent>
        <Typography variant="h4" component="h2">{`${owner?.name}'s Tisch`}</Typography>
        <Typography color="textSecondary">Created {table.created.toLocaleString()}</Typography>
        <List>
          {table.players.map((player) => (
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
        <GrowDiv />
        {canContinueTable(table) && (
          <Button variant="contained" color="primary" size="small">
            Weiter spielen
          </Button>
        )}
        {!started && (
          <Button
            variant="contained"
            color="primary"
            disabled={!canStartTable(table)}
            onClick={() => dispatch(startTable(table.id))}
          >
            Starten
          </Button>
        )}{' '}
      </CardActions>
    </Card>
  );
}
