import * as React from 'react';
import { TableState } from 'model/table';
import { Card, toUserMap } from 'model/core';
import MatchView from './MatchView';
import { toPlayerMap } from 'model/players';
import { makeStyles } from '@material-ui/core';

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
  table: TableState;
  playCard: (card: Card) => void;
}

export default function TableView({ table: { match, table }, playCard }: Props) {
  const classes = useStyles();
  if (!match) {
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
