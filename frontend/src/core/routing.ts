import { createSlice, PayloadAction } from "@reduxjs/toolkit";
import { RootState } from "../app/store";

export type Location = "login" | "room" | "table";
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
});
export const selectLocation = (root: RootState) => root.core.routing.location;
export const { setLocation } = locationSlice.actions;
export default locationSlice.reducer;
