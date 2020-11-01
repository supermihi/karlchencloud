import { CaseReducer, createSlice, Draft, PayloadAction } from '@reduxjs/toolkit';
import { Card } from 'model/core';
import { Match, newGame, PlayedCard } from 'model/match';
import { ActionKind, AsyncState, createGameThunk } from './asyncs';
import { newDeclareRequest, newPlayCardRequest } from 'api/modelToPb';
import * as api from 'api/karlchen_pb';
import { getPosition, nextPos, Pos } from 'model/players';
import { toDeclareResult } from 'model/apiconv';
import * as events from 'app/session/events';
import { startTable } from './table';
import { DeclareResult } from 'model/auction';
import { selectCurrentTableOrThrow, selectPlayers } from './selectors';
import { selectAuthenticatedClientOrThrow } from 'app/session';

export const playCard = createGameThunk<Card, PlayedCard>(
  ActionKind.playCard,
  async (card, { client: { client, meta }, getState }) => {
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
export interface CurrentMatchState extends AsyncState {
  match: Match | null;
}

const initialState: CurrentMatchState = {
  match: null,
};

const reducePlayedCard: CaseReducer<CurrentMatchState, PayloadAction<PlayedCard>> = (
  state,
  { payload }
) => {
  if (state.match === null) {
    return;
  }
  const trick = state.match.game.currentTrick;
  trick.cards.push(payload.card);
  trick.winner = payload.trickWinner;
  if (payload.player === Pos.bottom) {
    const card = state.match.cards.findIndex(
      (c) => c.rank === payload.card.rank && c.suit === payload.card.suit
    );
    state.match.cards.splice(card);
  }
  state.match.turn = nextPos(payload.player);
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
      });
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
