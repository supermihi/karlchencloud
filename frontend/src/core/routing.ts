import { createSlice, PayloadAction } from "@reduxjs/toolkit";
import { RootState } from "../app/store";
import { tryLogin, register } from "core/session/slice";

export type Location = "login" | "lobby" | "table";
export interface RoutingState {
  location: Location;
}
const initialState: RoutingState = { location: "login" };
export const locationSlice = createSlice({
  initialState,
  name: "location",
  reducers: {
    setLocation: (state, { payload }: PayloadAction<Location>) => {
      state.location = payload;
    },
  },
  extraReducers: (builder) =>
    builder
      .addCase(tryLogin.fulfilled.type, (state) => {
        state.location = "lobby";
      })
      .addCase(register.fulfilled.type, (state) => {
        state.location = "lobby";
      }),
});
export const selectLocation = (root: RootState) => root.core.routing.location;
export const { setLocation } = locationSlice.actions;
export default locationSlice.reducer;
