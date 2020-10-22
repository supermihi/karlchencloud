import * as React from 'react';
import { Table } from 'model/table';
import { Card, toUserMap } from 'model/core';
import MatchView from './MatchView';
import { toPlayerMap } from 'model/players';
import { makeStyles } from '@material-ui/core';
import { Match } from 'model/match';
import { TablePhase } from 'api/karlchen_pb';

const useStyles = makeStyles((theme) => ({
  main: {
    paddingTop: theme.spacing(2),
    paddingBottom: theme.spacing(2),
    boxSizing: 'border-box',
    width: '100%',
    height: '100%',
  },
}));
interface Props {
  table: Table;
  match: Match;
  playCard: (card: Card) => void;
}

export default function TableView({ table, match, playCard }: Props) {
  const classes = useStyles();
  if (table.phase !== TablePhase.PLAYING) {
    return <span className={classes.main}>'no match'</span>;
  }

  const users = toUserMap(table.players);
  const players = toPlayerMap(match.players, users);
  return (
    <div className={classes.main}>
      <MatchView match={match} players={players} playCard={playCard} />
    </div>
  );
}
