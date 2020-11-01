import React from 'react';
import { User } from 'model/core';
import { BidType, GameType } from 'api/karlchen_pb';
import HourglassEmptyIcon from '@material-ui/icons/HourglassEmpty';
import OfflineBoltIcon from '@material-ui/icons/OfflineBolt';
import { gameTypeString, bidString } from 'api/helpers';
import Typography from '@material-ui/core/Typography';
import Paper from '@material-ui/core/Paper';
import makeStyles from '@material-ui/core/styles/makeStyles';
import { Declaration } from 'model/auction';

export interface Props {
  user: User;
  declaration?: Declaration;
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
  declaration,
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
      {declaration !== undefined && (
        <Typography variant="subtitle2">
          {declaration === Declaration.vorbehalt ? 'vorbehalt' : 'gesund'}
        </Typography>
      )}

      {bids.length > 0 && <Typography variant="body2">{bids.map(bidString).join(', ')}</Typography>}
    </Paper>
  );
}
