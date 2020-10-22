import { Action, configureStore, ThunkAction, ThunkDispatch } from '@reduxjs/toolkit';
import authReducer, { AuthState } from './auth/slice';
import sessionReducer, { SessionState } from './session';
import lobbyReducer, { LobbyState } from 'features/lobby/slice';
import gameReducer, { GameState } from './game';

export const store = configureStore<RootState>({
  reducer: {
    auth: authReducer,
    session: sessionReducer,
    lobby: lobbyReducer,
    game: gameReducer,
  },
});
export type RootState = {
  auth: AuthState;
  session: SessionState;
  lobby: LobbyState;
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
