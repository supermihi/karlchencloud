import { Action, combineReducers, ThunkAction, ThunkDispatch } from '@reduxjs/toolkit';
import { GameState, reducer as gameReducer } from 'game/state';
import { SessionState } from 'session/state';
import { reducer as sessionReducer } from 'session/slice';

export type RootState = {
  session: SessionState;
  game: GameState;
};

export const reducer = combineReducers({
  session: sessionReducer,
  game: gameReducer,
});

export type AsyncThunkConfig = { state: RootState };
export type AppThunk<ReturnType = void> = ThunkAction<
  ReturnType,
  RootState,
  unknown,
  Action<string>
>;
export type AppDispatch = ThunkDispatch<RootState, {}, Action>;
export type Dispatchable =
  | Action<string>
  | ThunkAction<unknown, RootState, unknown, Action<string>>;
