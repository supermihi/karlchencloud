import { User } from "model/core";
import { createSlice, PayloadAction } from "@reduxjs/toolkit";

export interface UsersState {
  users: User[];
}

const initialState: UsersState = { users: [] };
const usersSlice = createSlice({
  name: "users",
  initialState,
  reducers: {
    add: (state, action: PayloadAction<User>) => {
      state.users.push(action.payload);
    },
    set: (state, { payload: user }: PayloadAction<User>) => {
      const exist = state.users.find((u) => u.id === user.id);
      if (exist) {
        exist.name = user.name;
      } else {
        state.users.push(user);
      }
    },
  },
});
export const { add, set } = usersSlice.actions;

export default usersSlice.reducer;
