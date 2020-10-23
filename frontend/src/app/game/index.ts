import { combineReducers } from '@reduxjs/toolkit';
import type { RootState } from 'app/store';
import { reducer as matchReducer, CurrentMatchState } from './match';
import { reducer as tableReducer, TableState } from './table';

export const selectGame = (state: RootState) => state.game;
export type GameState = {
  table: TableState;
  match: CurrentMatchState;
};
export default combineReducers<GameState>({
  table: tableReducer,
  match: matchReducer,
});
