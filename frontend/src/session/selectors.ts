import { createSelector } from '@reduxjs/toolkit';
import { getAuthenticatedClient } from 'api/client';
import { RootState } from 'state';
import { SessionState } from './state';

export const selectSession = (state: RootState): SessionState => state.session;
export const selectClient = createSelector(selectSession, ({ activeSession }) =>
  activeSession !== null ? getAuthenticatedClient(activeSession.token) : null
);
export const selectAuthenticatedClientOrThrow = createSelector(selectClient, (client) => {
  if (!client) {
    throw new Error('not authenticated');
  }
  return client;
});
