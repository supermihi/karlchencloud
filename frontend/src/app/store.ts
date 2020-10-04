import { Action, configureStore, ThunkAction, ThunkDispatch } from '@reduxjs/toolkit';
import authReducer from './auth/slice';
import sessionReducer from './session';
import lobbyReducer from 'features/lobby/slice';
import gameReducer from './game';

export const store = configureStore({
  reducer: {
    auth: authReducer,
    session: sessionReducer,
    lobby: lobbyReducer,
    game: gameReducer,
  },
});

export type RootState = ReturnType<typeof store.getState>;
export type AsyncThunkConfig = { state: RootState };
export type AppThunk<ReturnType = void> = ThunkAction<
  ReturnType,
  RootState,
  unknown,
  Action<string>
>;
export type AppDispatch = ThunkDispatch<RootState, {}, Action>;
