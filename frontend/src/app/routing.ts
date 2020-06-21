import { RootState } from "./store";
import { selectAuth } from "app/auth/slice";
import { selectSession } from "./session";

export enum Location {
  register,
  login,
  lobby,
  table,
}
export function selectLocation(state: RootState): Location {
  if (selectSession(state).session) {
    return Location.lobby;
  }
  return selectAuth(state).storedLogin ? Location.login : Location.register;
}
