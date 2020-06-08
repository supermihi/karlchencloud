import React, { useCallback } from "react";
import RoomView from "./RoomView";
import { useSelector, useDispatch } from "react-redux";
import { selectRoomState, createTable, fetchTables } from "./slice";
import CircularProgress from "@material-ui/core/CircularProgress";
import Alert from "@material-ui/lab/Alert";
import { formatError } from "api/client";

export default () => {
  const state = useSelector(selectRoomState);
  const dispatch = useDispatch();
  const dispatchCreateTable = useCallback(() => dispatch(createTable()), [
    dispatch,
  ]);
  if (!state.loaded) {
    if (state.loading) {
      return <CircularProgress />;
    }
    if (state.error) {
      return <Alert>Error loading tables: {formatError(state.error)}</Alert>;
    }
    dispatch(fetchTables());
    return null;
  }
  return <RoomView tables={state.tables} createTable={dispatchCreateTable} />;
};
