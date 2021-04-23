import { combineReducers } from '@reduxjs/toolkit';
import { Match } from 'model/match';
import { Table } from 'model/table';
import { ActionState } from './playActions';
import { reducer as matchReducer } from './match';
import { reducer as actionReducer } from './playActions';
import { reducer as tableReducer } from './table';
export type PlayState = {
  match: Match | null;
  table: Table | null;
  action: ActionState;
};

export const reducer = combineReducers<PlayState>({
  match: matchReducer,
  table: tableReducer,
  action: actionReducer,
});
