import { User } from "model/core";
import { createSlice, PayloadAction } from "@reduxjs/toolkit";
import { RootState } from "../app/store";

export interface UsersState {
  byId: { [id: string]: User };
}

const initialState: UsersState = { byId: {} };
const usersSlice = createSlice({
  name: "users",
  initialState,
  reducers: {
    set: (state, { payload: user }: PayloadAction<User>) => {
      state.byId[user.id] = user;
    },
  },
});
export const { set } = usersSlice.actions;
export const selectUser = (id: string) => (state: RootState) =>
  state.core.users.byId[id];

export default usersSlice.reducer;
