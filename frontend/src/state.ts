import {
  Action,
  combineReducers,
  configureStore,
  ThunkAction,
  ThunkDispatch,
} from '@reduxjs/toolkit';
import { PlayState, reducer as playReducer } from 'play/state';
import { SessionState } from 'session/state';
import { reducer as sessionReducer } from 'session/slice';

export type RootState = {
  session: SessionState;
  play: PlayState;
};

export const reducer = combineReducers({
  session: sessionReducer,
  play: playReducer,
});
export const store = configureStore<RootState>({
  reducer,
});

export type AsyncThunkConfig = { state: RootState };
export type AppThunk<ReturnType = void> = ThunkAction<
  ReturnType,
  RootState,
  unknown,
  Action<string>
>;
export type AppDispatch = ThunkDispatch<RootState, unknown, Action>;
export type Dispatchable =
  | Action<string>
  | ThunkAction<unknown, RootState, unknown, Action<string>>;
