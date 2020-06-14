import React from "react";
import Lobby from "./Lobby";
import { useSelector, useDispatch } from "react-redux";
import { selectRoomState, createTable } from "./slice";
import CircularProgress from "@material-ui/core/CircularProgress";
import Alert from "@material-ui/lab/Alert";
import { formatError } from "api/client";

export default () => {
  const state = useSelector(selectRoomState);
  const dispatch = useDispatch();
  if (!state.loaded) {
    if (state.loading) {
      return <CircularProgress />;
    }
    if (state.error) {
      return <Alert>{formatError(state.error)}</Alert>;
    }
    return null;
  }
  return (
    <Lobby
      activeTable={state.activeTable}
      createTable={() => dispatch(createTable())}
    />
  );
};
