import React from 'react';
import MatchView from 'features/table/MatchView';
import { Match, Game } from 'model/match';
import { MatchPhase, GameType } from 'api/karlchen_pb';
import * as data from './mocks';

export default {
  title: 'Match/Match',
  component: MatchView,
};

const mode = {
  type: GameType.NORMAL_GAME,
  forehand: 'pl1',
};
const game: Game = {
  bids: {},
  completedTricks: 0,
  mode,
  currentTrick: data.trick(1, 4),
};

const match: Match = {
  phase: MatchPhase.GAME,
  players: data.players,
  cards: data.fullHand,
  details: game,
};
export const MatchTable = () => (
  <MatchView style={{ width: '80vw', height: '95vh' }} match={match} users={data.userMap} />
);
