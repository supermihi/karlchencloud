import { Draft } from '@reduxjs/toolkit';

export enum ActionKind {
  noAction = 'noAction',
  joinTable = 'joinTable',
  createTable = 'createTable',
  startTable = 'startTable',
  playCard = 'playCard',
  placeBid = 'placeBid',
  declare = 'declare',
}

export interface ActionError {
  action: ActionKind;
  error: any;
}

export interface AsyncState {
  pendingAction?: ActionKind;
  error?: ActionError;
}
export function clearPendingAndError(state: Draft<AsyncState>) {
  state.pendingAction = ActionKind.noAction;
  state.error = undefined;
}
