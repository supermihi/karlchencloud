import { createAsyncThunk, createSlice, Draft } from '@reduxjs/toolkit';
import { Card } from 'model/core';
import { dummyGame, Match, newGame } from 'model/match';
import { ActionError, ActionKind } from './asyncs';
import { createGameThunk } from './constants';
import { newDeclareRequest, newPlayCardRequest } from 'api/modelToPb';
import * as api from 'api/karlchen_pb';
import { newPlayerMap, nextPos } from 'model/players';
import { toDeclareResult } from 'model/apiconv';
import * as events from 'app/session/events';
import { startTable } from './table';
import { DeclareResult, emptyAuction } from 'model/auction';
import { AsyncThunkConfig } from 'app/store';
import { selectCurrentMatchOrThrow, selectCurrentTableOrThrow } from './selectors';
import { selectAuthenticatedClientOrThrow } from 'app/session';

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

export interface MatchState extends Match {
  pendingAction?: ActionKind;
  error?: ActionError;
}

export const declare = createAsyncThunk<
  DeclareResult & { gametype: api.GameType },
  api.GameType,
  AsyncThunkConfig
>('game/dispatch', async (gametype: api.GameType, thunkAPI) => {
  const state = thunkAPI.getState();
  const table = selectCurrentTableOrThrow(state);
  const match = selectCurrentMatchOrThrow(state);
  const req = newDeclareRequest(gametype, table.id);
  const { client, meta } = selectAuthenticatedClientOrThrow(state);
  const ans = await client.declare(req, meta);
  return { ...toDeclareResult(ans, match.players), gametype };
});

const initialState: MatchState = {
  players: newPlayerMap((_) => ''),
  phase: api.MatchPhase.AUCTION,
  cards: [],
  auction: emptyAuction(),
  game: dummyGame(),
};

const matchSlice = createSlice({
  name: 'game/match',
  initialState: initialState,
  reducers: {},
  extraReducers: (builder) => {
    builder
      .addCase(events.matchStarted, (_, { payload }) => payload)
      .addCase(startTable.fulfilled, (_, { payload }) => payload)
      .addCase(events.tableChanged, (_, { payload }) => payload?.match)
      .addCase(playCard.fulfilled, (match, { payload }) => {
        const card = match.cards.findIndex(
          (c) => c.rank === payload.rank && c.suit === payload.suit
        );
        match.cards.splice(card);
      })
      .addCase(events.playerDeclared, (match, { payload }) => {
        reduceDeclaration(match, payload);
      })
      .addCase(declare.fulfilled, (match, { payload }) => {
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
