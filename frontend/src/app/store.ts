import { configureStore } from "@reduxjs/toolkit";
import loginReducer from "core/login";
import usersReducer from "core/users";

import roomReducer from "features/room/slice";

export const store = configureStore({
  reducer: {
    login: loginReducer,
    users: usersReducer,
    room: roomReducer,
  },
});

export type RootState = ReturnType<typeof store.getState>;
export type AsyncThunkConfig = { state: RootState };
