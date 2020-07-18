import React from 'react';
import Grid from '@material-ui/core/Grid';
import PlayerView from './PlayerView';
import { User } from 'model/core';

interface Props {
  left: User;
  top: User;
  right: User;
}
export default function PlayersView({ left, top, right }: Props) {
  return (
    <Grid container spacing={3}>
      <Grid item xs={4}>
        <PlayerView user={left} />
      </Grid>
      <Grid item xs={4}>
        <PlayerView user={top} />
      </Grid>
      <Grid item xs={4}>
        <PlayerView user={right} />
      </Grid>
    </Grid>
  );
}
