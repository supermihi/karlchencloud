import { createAction } from '@reduxjs/toolkit';
import { DeclareResult } from 'model/auction';
import { User } from 'model/core';
import { Match } from 'model/match';
import { TableState } from 'model/table';

export const sessionStarted = createAction<string>('sessionStarted');
export const tableChanged = createAction<TableState | null>('tableChanged');
export const memberJoined = createAction<User>('memberJoined');
export const memberLeft = createAction<string>('memberLeft');
export const memberStatusChanged = createAction<User>('memberStatusChanged');
export const matchStarted = createAction<Match>('matchStarted');
export const playerDeclared = createAction<DeclareResult>('playerDeclared');
