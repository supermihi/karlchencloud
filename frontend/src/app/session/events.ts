import { createAction } from '@reduxjs/toolkit';
import { DeclareResult } from 'model/auction';
import { User } from 'model/core';
import { Match } from 'model/match';
import { TableState } from 'model/table';

export const sessionStarted = createAction<string>('event/sessionStarted');
export const tableChanged = createAction<TableState | null>('event/tableChanged');
export const memberJoined = createAction<User>('event/memberJoined');
export const memberLeft = createAction<string>('event/memberLeft');
export const memberStatusChanged = createAction<User>('event/memberStatusChanged');
export const matchStarted = createAction<Match>('event/matchStarted');
export const playerDeclared = createAction<DeclareResult>('event/playerDeclared');
