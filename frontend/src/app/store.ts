import { Action, configureStore, ThunkAction } from "@reduxjs/toolkit";
import authReducer from "./auth/slice";
import sessionReducer from "./session";

export const store = configureStore({
  reducer: {
    auth: authReducer,
    session: sessionReducer,
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
