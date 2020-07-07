import { TableState } from 'model/table';

export enum ActionKind {
  noAction = 'noAction',
  joinTable = 'joinTable',
  createTable = 'createTable',
  startTable = 'startTable',
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
