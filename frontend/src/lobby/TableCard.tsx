import * as React from 'react';

import Card from '@material-ui/core/Card';
import CardContent from '@material-ui/core/CardContent';
import CardHeader from '@material-ui/core/CardHeader';
import Button from '@material-ui/core/Button';
import PlayIcon from '@material-ui/icons/PlayCircleFilled';
import makeStyles from '@material-ui/core/styles/makeStyles';
import Typography from '@material-ui/core/Typography';
import IconButton from '@material-ui/core/IconButton';
import { canStartTable, waitingForPlayers, Table } from 'model/table';
import { TablePhase } from 'api/karlchen_pb';
import InviteDialog from './InviteLinkDialog';
import PlayerItem from './PlayerItem';
import { User } from 'model/core';
import { createStyles } from '@material-ui/core/styles';
import EmailOutlined from '@material-ui/icons/EmailOutlined';

interface Props {
  table: Table;
  me: User;
  startTable: () => void;
}

const useStyles = makeStyles((theme) =>
  createStyles({
    continueButton: {
      marginLeft: 'auto',
    },
    note: {
      color: theme.palette.text.hint,
    },
    chips: {
      display: 'flex',
      flexWrap: 'wrap',
      '& > *': {
        margin: theme.spacing(0.5),
      },
    },
    grow: { flexGrow: 1 },
  })
);

export default function CurrentTableView({ table, me, startTable }: Props): React.ReactElement {
  const classes = useStyles();
  const [inviteOpen, setInviteOpen] = React.useState(false);

  const started = table.phase !== TablePhase.NOT_STARTED;
  const owner = table.members.find((p) => p.id === table.owner);

  return (
    <Card>
      <CardHeader
        title={`${owner?.name}'s Tisch`}
        subheader={`Erstellt ${table.created.toLocaleString()}`}
        action={
          <IconButton color="primary" disabled={!canStartTable(table, me)} onClick={startTable}>
            <PlayIcon style={{ fontSize: '2.5rem' }} />
          </IconButton>
        }
      />
      <CardContent>
        <div className={classes.chips}>
          <Typography component="span">Spieler:</Typography>
          {table.members.map((player) => (
            <PlayerItem me={player.id === me.id} player={player} key={player.id} />
          ))}
          <div className={classes.grow} />
          {!started && (
            <>
              <Button onClick={() => setInviteOpen(true)} startIcon={<EmailOutlined />}>
                einladen
              </Button>
              <InviteDialog
                open={inviteOpen}
                handleClose={() => setInviteOpen(false)}
                inviteCode={table.invite || ''}
              />
            </>
          )}
        </div>
        {waitingForPlayers(table) && (
          <Typography variant="body1" className={classes.note}>
            Warte auf Mitspieler …
          </Typography>
        )}
      </CardContent>
    </Card>
  );
}
