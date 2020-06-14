import authReducer from "../features/auth/slice";
import usersReducer from "./users";
import routingReducer from "./routing";
import { combineReducers } from "redux";
export default combineReducers({
  auth: authReducer,
  users: usersReducer,
  routing: routingReducer,
});
