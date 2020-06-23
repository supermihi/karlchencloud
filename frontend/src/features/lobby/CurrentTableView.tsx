import * as React from 'react';

import Card from '@material-ui/core/Card';
import CardContent from '@material-ui/core/CardContent';
import CardActions from '@material-ui/core/CardActions';
import Button from '@material-ui/core/Button';
import List from '@material-ui/core/List';
import ListItem from '@material-ui/core/ListItem';
import ListItemText from '@material-ui/core/ListItemText';
import makeStyles from '@material-ui/core/styles/makeStyles';
import ListItemAvatar from '@material-ui/core/ListItemAvatar';
import Avatar from '@material-ui/core/Avatar';
import Typography from '@material-ui/core/Typography';

import { canStartTable, TableState } from 'model/table';
import { TablePhase } from 'api/karlchen_pb';
import GrowDiv from 'components/GrowDiv';
import InviteDialog from './InviteDialog';

interface Props {
  table: TableState;
}
const useStyles = makeStyles((theme) => ({
  continueButton: {
    marginLeft: 'auto',
  } as const,
  online: {
    color: theme.palette.success.main,
  } as const,
}));

export default function CurrentTableView({ table }: Props) {
  const { table: data, phase } = table;
  const owner = data.players.find((p) => p.id === data.owner);
  const classes = useStyles();
  const [inviteOpen, setInviteOpen] = React.useState(false);
  return (
    <Card>
      <CardContent>
        <Typography
          variant="h4"
          component="h2"
        >{`${owner?.name}'s Tisch`}</Typography>
        <Typography color="textSecondary">
          Created {data.created.toLocaleString()}
        </Typography>
        <List>
          {data.players.map((player) => (
            <ListItem key={player.id}>
              <ListItemAvatar>
                <Avatar>{player.name[0].toUpperCase()}</Avatar>
              </ListItemAvatar>
              <ListItemText
                primary={player.name}
                secondary={
                  player.online && <em className={classes.online}>online</em>
                }
              />
            </ListItem>
          ))}
        </List>
      </CardContent>
      <CardActions>
        {phase === TablePhase.NOT_STARTED && (
          <>
            <Button onClick={() => setInviteOpen(true)}>
              Mitspieler einladen
            </Button>
            <InviteDialog
              open={inviteOpen}
              handleClose={() => setInviteOpen(false)}
              inviteCode={table.table.invite || ''}
            />
          </>
        )}
        <GrowDiv />
        {(phase === TablePhase.PLAYING ||
          phase === TablePhase.BETWEEN_GAMES) && (
          <Button variant="contained" color="primary" size="small">
            Weiter spielen
          </Button>
        )}
        {phase === TablePhase.NOT_STARTED && (
          <Button
            variant="contained"
            color="primary"
            disabled={!canStartTable(table)}
          >
            Starten
          </Button>
        )}{' '}
      </CardActions>
    </Card>
  );
}
