import { createSlice, Draft } from '@reduxjs/toolkit';
import { Card } from 'model/core';
import { Match, newGame } from 'model/match';
import { ActionKind, AsyncState, createGameThunk } from './asyncs';
import { newDeclareRequest, newPlayCardRequest } from 'api/modelToPb';
import * as api from 'api/karlchen_pb';
import { nextPos } from 'model/players';
import { toDeclareResult } from 'model/apiconv';
import * as events from 'app/session/events';
import { startTable } from './table';
import { DeclareResult } from 'model/auction';
import { selectCurrentTableOrThrow, selectPlayers } from './selectors';
import { selectAuthenticatedClientOrThrow } from 'app/session';

export const playCard = createGameThunk(
  ActionKind.playCard,
  async (card: Card, { client: { client, meta }, getState }) => {
    const tableId = selectCurrentTableOrThrow(getState()).id;
    const req = newPlayCardRequest(card, tableId);
    await client.playCard(req, meta);
    return card;
  }
);

export interface CurrentMatchState extends AsyncState {
  match: Match | null;
}

export const declare = createGameThunk<api.GameType, DeclareResult & { gametype: api.GameType }>(
  ActionKind.declare,
  async (gametype: api.GameType, thunkAPI) => {
    const state = thunkAPI.getState();
    const table = selectCurrentTableOrThrow(state);
    const players = selectPlayers(state);
    const req = newDeclareRequest(gametype, table.id);
    const { client, meta } = selectAuthenticatedClientOrThrow(state);
    const ans = await client.declare(req, meta);
    return { ...toDeclareResult(ans, players), gametype };
  }
);

const initialState: CurrentMatchState = {
  match: null,
};

const matchSlice = createSlice({
  name: 'game/match',
  initialState: initialState,
  reducers: {},
  extraReducers: (builder) => {
    builder
      .addCase(events.matchStarted, (_, { payload }) => ({ match: payload }))
      .addCase(startTable.fulfilled, (_, { payload }) => ({ match: payload }))
      .addCase(events.tableChanged, (_, { payload }) => ({ match: payload?.match ?? null }))
      .addCase(playCard.fulfilled, ({ match }, { payload }) => {
        if (match === null) return;
        const card = match.cards.findIndex(
          (c) => c.rank === payload.rank && c.suit === payload.suit
        );
        match.cards.splice(card);
      })
      .addCase(events.playerDeclared, ({ match }, { payload }) => {
        if (match === null) return;
        reduceDeclaration(match, payload);
      })
      .addCase(declare.fulfilled, ({ match }, { payload }) => {
        if (match === null) return;
        reduceDeclaration(match, payload);
        match.auction.ownDeclaration = payload.gametype;
      });
  },
});

function reduceDeclaration(match: Draft<Match>, { declaration, mode, player }: DeclareResult) {
  match.auction.declarations[player] = declaration;
  if (mode !== null) {
    match.phase = api.MatchPhase.GAME;
    match.game = newGame(mode);
  } else if (match.turn !== undefined) {
    match.turn = nextPos(match.turn);
  }
}
export const { actions, reducer } = matchSlice;
