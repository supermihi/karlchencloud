import type { RootState } from 'state';
import { CaseReducer, createSlice, Draft, PayloadAction } from '@reduxjs/toolkit';
import { Card, User } from 'model/core';
import { Match, newGame, PlayedCard } from 'model/match';
import { ActionKind, createGameThunk } from './asyncs';
import { newDeclareRequest, newPlayCardRequest } from 'api/modelToPb';
import * as api from 'api/karlchen_pb';
import { getPosition, nextPos, Pos } from 'model/players';
import { toDeclareResult, toMatch } from 'model/apiconv';
import { DeclareResult } from 'model/auction';
import { selectCurrentTableOrThrow, selectPlayers } from './selectors';
import { Table } from 'model/table';
import { selectAuthenticatedClientOrThrow } from 'session/selectors';
import { tableId } from 'api/helpers';
import * as events from 'session/events';
import { createTable, joinTable } from './thunks';

export const selectGame = (state: RootState) => state.game;

export type GameState = {
  match: Match | null;
  table: Table | null;
  loading: boolean;
  error?: unknown;
};

export const startTable = createGameThunk<string, Match>(
  ActionKind.startTable,
  async (id, { client, meta }) => {
    const match = await client.startTable(tableId(id), meta);
    return toMatch(match);
  }
);

export const playCard = createGameThunk<Card, PlayedCard>(
  ActionKind.playCard,
  async (card, { client, meta, getState }) => {
    const tableId = selectCurrentTableOrThrow(getState()).id;
    const req = newPlayCardRequest(card, tableId);
    const result = await client.playCard(req, meta);
    let winner: Pos | undefined = undefined;
    if (result.hasTrickWinner()) {
      const winnerId = (result.getTrickWinner() as api.PlayerValue).getUserId();
      const players = selectPlayers(getState());
      winner = getPosition(players, winnerId);
    }
    return { card, player: Pos.bottom, trickWinner: winner };
  }
);

export const declare = createGameThunk<api.GameType, DeclareResult & { gametype: api.GameType }>(
  ActionKind.declare,
  async (gametype: api.GameType, thunkAPI) => {
    const state = thunkAPI.getState();
    const table = selectCurrentTableOrThrow(state);
    const req = newDeclareRequest(gametype, table.id);
    const { client, meta } = selectAuthenticatedClientOrThrow(state);
    const ans = await client.declare(req, meta);
    const players = selectPlayers(state);
    return { ...toDeclareResult(ans, players), gametype };
  }
);

const initialState: GameState = {
  match: null,
  table: null,
  loading: false,
};

const reducePlayedCard: CaseReducer<GameState, PayloadAction<PlayedCard>> = (
  state,
  { payload: { card, trickWinner, player } }
) => {
  if (state.match === null) {
    return;
  }
  const trick = state.match.game.currentTrick;
  trick.cards.push(card);
  trick.winner = trickWinner;
  if (player === Pos.bottom) {
    const cardIndex = state.match.cards.findIndex(
      (c) => c.rank === card.rank && c.suit === card.suit
    );
    state.match.cards.splice(cardIndex, 1);
  }
  state.match.turn = nextPos(player);
};

const matchSlice = createSlice({
  name: 'game/match',
  initialState: initialState,
  reducers: {},
  extraReducers: (builder) => {
    builder
      .addCase(events.matchStarted, (state, { payload }) => {
        state.match = payload;
      })
      .addCase(startTable.fulfilled, (state, { payload }) => {
        state.match = payload;
      })
      .addCase(events.tableChanged, (state, { payload }) => {
        state.match = payload?.match ?? null;
      })
      .addCase(playCard.fulfilled, reducePlayedCard)
      .addCase(events.cardPlayed, reducePlayedCard)
      .addCase(events.playerDeclared, ({ match }, { payload }) => {
        if (match === null) return;
        reduceDeclaration(match, payload);
      })
      .addCase(declare.fulfilled, ({ match }, { payload }) => {
        if (match === null) return;
        reduceDeclaration(match, payload);
        match.auction.ownDeclaration = payload.gametype;
      })
      .addCase(events.tableChanged, (state, { payload }) => ({
        ...state,
        table: payload?.table ?? null,
      }))
      .addCase(events.memberJoined, ({ table }, { payload }) => {
        table?.players.push(payload);
      })
      .addCase(events.memberLeft, ({ table }, { payload: id }: PayloadAction<string>) => {
        if (table === null) return;
        const index = table.players.findIndex((p) => p.id === id);
        if (index !== -1) {
          table.players.splice(index, 1);
        }
      })
      .addCase(events.memberStatusChanged, ({ table }, { payload: user }: PayloadAction<User>) => {
        if (table === null) return;
        const index = table.players.findIndex((p) => p.id === user.id);
        if (index !== -1) {
          table.players.splice(index, 1, user);
        }
      })
      .addCase(events.matchStarted, ({ table }) => {
        if (table === null) return;
        table.phase = api.TablePhase.PLAYING;
      })
      .addCase(startTable.fulfilled, (state, { payload }) => {
        if (state.table === null) return;
        state.table.phase = api.TablePhase.PLAYING;
        state.match = payload;
      })
      .addCase(createTable.fulfilled, (state, { payload: table }) => {
        state.table = table;
        state.match = null;
      })
      .addCase(joinTable.fulfilled, (state, { payload: table }) => (state.table = table));
  },
});

function reduceDeclaration(match: Draft<Match>, { declaration, mode, player }: DeclareResult) {
  match.auction.declarations[player] = declaration;
  if (mode !== null) {
    match.phase = api.MatchPhase.GAME;
    match.game = newGame(mode);
    match.turn = mode.forehand;
  } else {
    match.turn = nextPos(player);
  }
}

export const { actions, reducer } = matchSlice;
