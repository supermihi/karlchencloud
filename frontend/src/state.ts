import { Action, ThunkAction, ThunkDispatch } from '@reduxjs/toolkit';
import { GameState } from 'game';
import { SessionState } from 'session/state';

export type RootState = {
  session: SessionState;
  game: GameState;
};
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
