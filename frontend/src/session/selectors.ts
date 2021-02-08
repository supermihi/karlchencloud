import { createSelector } from '@reduxjs/toolkit';
import { getAuthenticatedClient } from 'api/client';
import { RootState } from 'state';

export const selectSession = (state: RootState) => state.session;
export const selectClient = createSelector(selectSession, ({ session }) =>
  session !== null ? getAuthenticatedClient(session.id, session.secret) : null
);
export const selectAuthenticatedClientOrThrow = createSelector(selectClient, (client) => {
  if (!client) {
    throw new Error('not authenticated');
  }
  return client;
});
