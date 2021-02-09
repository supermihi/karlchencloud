import { createAction } from '@reduxjs/toolkit';
import { TableState } from 'model/apiconv';
import { DeclareResult } from 'model/auction';
import { User } from 'model/core';
import { Match, PlayedCard } from 'model/match';

export const sessionStarted = createAction<string>('event/sessionStarted');
export const tableChanged = createAction<TableState | null>('event/tableChanged');
export const memberJoined = createAction<User>('event/memberJoined');
export const memberLeft = createAction<string>('event/memberLeft');
export const memberStatusChanged = createAction<User>('event/memberStatusChanged');
export const matchStarted = createAction<Match>('event/matchStarted');
export const playerDeclared = createAction<DeclareResult>('event/playerDeclared');
export const cardPlayed = createAction<PlayedCard>('event/cardPlayed');
export const notImplementedEvent = createAction<string>('event/unimplemented');
