import { TableState } from 'model/table';

export enum ActionKind {
  noAction = 0,
  joinTable = 1,
  createTable = 2,
  startTable = 3,
}

export interface ActionError {
  action: ActionKind;
  error: any;
}
export interface GameState {
  currentTable: TableState | null;
  pendingAction: ActionKind;
  error?: ActionError;
}

export const initialState: GameState = {
  currentTable: null,
  pendingAction: ActionKind.noAction,
};
