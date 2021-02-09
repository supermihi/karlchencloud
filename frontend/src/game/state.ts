import { combineReducers } from '@reduxjs/toolkit';
import { Match } from 'model/match';
import { Table } from 'model/table';
import { ActionState } from './gameActions';
import { reducer as matchReducer } from './match';
import { reducer as actionReducer } from './gameActions';
import { reducer as tableReducer } from './table';
export type GameState = {
  match: Match | null;
  table: Table | null;
  action: ActionState;
};

export const reducer = combineReducers<GameState>({
  match: matchReducer,
  table: tableReducer,
  action: actionReducer,
});
