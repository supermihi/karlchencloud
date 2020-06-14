import { Action, configureStore, ThunkAction } from "@reduxjs/toolkit";
import coreReducer from "core";
import roomReducer from "features/room/slice";

export const store = configureStore({
  reducer: {
    core: coreReducer,
    room: roomReducer,
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
