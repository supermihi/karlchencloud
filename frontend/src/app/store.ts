import {
  Action,
  configureStore,
  ThunkAction,
  getDefaultMiddleware,
} from "@reduxjs/toolkit";
import coreReducer from "core";
import lobbyReducer from "features/lobby/slice";
import createSagaMiddleware from "redux-saga";
import sagas from "./sagas";
const sagaMiddleware = createSagaMiddleware();

export const store = configureStore({
  reducer: {
    core: coreReducer,
    lobby: lobbyReducer,
  },
  middleware: [...getDefaultMiddleware(), sagaMiddleware],
});
sagaMiddleware.run(sagas);

export type RootState = ReturnType<typeof store.getState>;
export type AsyncThunkConfig = { state: RootState };
export type AppThunk<ReturnType = void> = ThunkAction<
  ReturnType,
  RootState,
  unknown,
  Action<string>
>;
