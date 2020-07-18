import * as React from 'react';
import Grid from '@material-ui/core/Grid';
import { TableState } from 'model/table';
import CardView from './CardsView';
import PlayersView from './PlayersView';
import { User } from 'model/core';
import TrickView from './TrickView';

interface Props {
  table: TableState;
}

export default function TableView({ table: { match, table } }: Props) {
  if (!match) {
    return <span>'no match'</span>;
  }
  const findUser = (id: string): User =>
    table.players.find((p) => p.id === id) || { id: 'none', name: 'error' };
  return (
    <div>
      <PlayersView
        left={findUser(match.players[1])}
        top={findUser(match.players[2])}
        right={findUser(match.players[3])}
      />
      <Grid container spacing={2} justify="center">
        <Grid item xs={12}>
          <div style={{ display: 'flex', justifyContent: 'space-around' }}>
            {match.cards.map((c) => (
              <CardView card={c} />
            ))}
          </div>
        </Grid>
      </Grid>
    </div>
  );
}
