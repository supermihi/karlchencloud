import { CaseReducer, createSlice, PayloadAction } from '@reduxjs/toolkit';
import { DeclareResult } from 'model/auction';
import { inAuction, inGame, Match, PlayedCard } from 'model/match';
import { afterDeclaration, afterPlayedCard } from 'model/mutations';
import update from 'immutability-helper';
import * as events from 'session/events';
import { createTable, declare, joinTable, playCard, startTable } from './thunks';

type MatchState = Match | null;
const initialState: MatchState = null;

const reducePlayedCard: CaseReducer<MatchState, PayloadAction<PlayedCard>> = (
  match,
  { payload }
) => {
  if (!match || !inGame(match)) {
    return;
  }
  return afterPlayedCard(match, payload);
};
const reduceDeclaration: CaseReducer<MatchState, PayloadAction<DeclareResult>> = (
  match,
  { payload }
) => {
  if (!match || !inAuction(match)) return null;
  return afterDeclaration(match, payload);
};

const matchSlice = createSlice({
  name: 'play/match',
  initialState: initialState as MatchState,
  reducers: {},
  extraReducers: (builder) => {
    builder
      .addCase(events.matchStarted, (_, { payload }) => payload)
      .addCase(startTable.fulfilled, (_, { payload }) => payload)
      .addCase(createTable.fulfilled, () => null)
      .addCase(joinTable.fulfilled, (_, { payload: { match } }) => match)
      .addCase(events.tableChanged, (_, { payload }) => payload?.match ?? null)
      .addCase(playCard.fulfilled, reducePlayedCard)
      .addCase(events.cardPlayed, reducePlayedCard)
      .addCase(events.playerDeclared, reduceDeclaration)
      .addCase(declare.fulfilled, (match, action) => {
        if (match === null) return match;
        const ans = reduceDeclaration(match, action);
        return update(ans, { auction: { ownDeclaration: { $set: action.payload.gametype } } });
      });
  },
});
export const { reducer } = matchSlice;
