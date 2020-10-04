import { TableState } from 'model/table';

export enum ActionKind {
  noAction = 'noAction',
  joinTable = 'joinTable',
  createTable = 'createTable',
  startTable = 'startTable',
  playCard = 'playCard',
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
