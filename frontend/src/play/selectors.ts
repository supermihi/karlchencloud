import { createSelector } from '@reduxjs/toolkit';
import { RootState } from 'state';
import { PlayState } from './state';

export const selectPlay = (state: RootState): PlayState => state.play;
export const selectTable = createSelector(selectPlay, (g) => g.table);
export const selectMatch = createSelector(selectPlay, (g) => g.match);

export const selectCurrentTableOrThrow = createSelector(selectTable, (t) => {
  if (t === null) {
    throw new Error('no current table');
  }
  return t;
});

export const selectCurrentMatchOrThrow = createSelector(selectMatch, (m) => {
  if (m === null) {
    throw new Error('no current match');
  }
  return m;
});

export const selectPlayers = createSelector(selectCurrentMatchOrThrow, (m) => m.players);
