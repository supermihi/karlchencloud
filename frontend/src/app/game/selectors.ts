import { createSelector } from '@reduxjs/toolkit';
import { selectGame } from '.';

export const selectTable = createSelector(selectGame, (g) => g.currentTable);
export const selectCurrentTableOrThrow = createSelector(selectTable, (t) => {
  if (t === null) {
    throw new Error('no current table');
  }
  return t;
});
