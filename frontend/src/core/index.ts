import sessionReducer from "./session/slice";
import usersReducer from "./users";
import routingReducer from "./routing";
import { combineReducers } from "redux";
export default combineReducers({
  session: sessionReducer,
  users: usersReducer,
  routing: routingReducer,
});
