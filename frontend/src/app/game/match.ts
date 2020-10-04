import { createSlice } from '@reduxjs/toolkit';
import { Card } from 'model/core';
import { Match } from 'model/match';
import { ActionKind } from './state';
import { createGameThunk } from './constants';
import { newPlayCardRequest } from 'api/modelToPb';

export const playCard = createGameThunk(
  ActionKind.playCard,
  async (card: Card, { client: { client, meta }, tableId }) => {
    if (tableId === undefined) {
      throw new Error('table ID not set');
    }
    const req = newPlayCardRequest(card, tableId);
    await client.playCard(req, meta);
    return card;
  }
);

const matchSlice = createSlice({
  name: 'game/table/match',
  initialState: {} as Match,
  reducers: {},
  extraReducers: (builder) => {
    builder.addCase(playCard.fulfilled, (match, { payload }) => {
      const card = match.cards.findIndex((c) => c.rank === payload.rank && c.suit === payload.suit);
      match.cards.splice(card);
    });
  },
});

export default matchSlice.reducer;
