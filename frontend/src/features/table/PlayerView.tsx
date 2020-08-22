import React from 'react';
import { User } from 'model/core';
import { BidType, GameType } from 'api/karlchen_pb';
import HourglassEmptyIcon from '@material-ui/icons/HourglassEmpty';
import OfflineBoltIcon from '@material-ui/icons/OfflineBolt';
import { gameTypeString, bidString } from 'api/helpers';
import Typography from '@material-ui/core/Typography';
import Paper from '@material-ui/core/Paper';
import makeStyles from '@material-ui/core/styles/makeStyles';

export interface Props {
  user: User;
  vorbehalt?: boolean;
  bids: BidType[];
  soloGame?: GameType;
  turn: boolean;
}

const useStyles = makeStyles((theme) => ({
  root: {
    padding: theme.spacing(1),
    backgroundColor: 'rgba(255,255,255,.5)',
  },
}));

export default function PlayerView({
  user: { name, online },
  soloGame,
  vorbehalt,
  bids,
  turn,
}: Props) {
  const classes = useStyles();
  return (
    <Paper variant="outlined" className={classes.root}>
      <Typography variant="h6">
        <span>
          {name}
          {turn && <HourglassEmptyIcon fontSize="inherit" />}
          {!online && <OfflineBoltIcon color="error" fontSize="inherit" />}
        </span>
      </Typography>
      {soloGame && <Typography variant="subtitle2">spielt {gameTypeString(soloGame)}</Typography>}
      {vorbehalt && <Typography variant="subtitle2">vorbehalt</Typography>}

      {bids.length > 0 && <Typography variant="body2">{bids.map(bidString).join(', ')}</Typography>}
    </Paper>
  );
}
