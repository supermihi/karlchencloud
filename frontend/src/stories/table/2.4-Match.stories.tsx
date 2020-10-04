import React from 'react';
import MatchView from 'features/table/MatchView';
import { Match, Game } from 'model/match';
import { MatchPhase, GameType, BidType } from 'api/karlchen_pb';
import * as data from './mocks';
import { Pos, toPlayerMap } from 'model/players';
import { action } from '@storybook/addon-actions';

export default {
  title: 'Match/Match',
  component: MatchView,
};

const mode = {
  type: GameType.NORMAL_GAME,
  forehand: 'pl1',
};
const game: Game = {
  bids: {
    [Pos.left]: [],
    [Pos.top]: [],
    [Pos.right]: [BidType.RE_BID, BidType.RE_NO_NINETY],
    [Pos.bottom]: [],
  },
  completedTricks: 0,
  mode,
  currentTrick: data.trick(Pos.left, 3),
};

const match: Match = {
  phase: MatchPhase.GAME,
  players: data.players,
  cards: data.fullHand,
  details: game,
  turn: Pos.bottom,
};
export const MatchTable = () => {
  return (
    <MatchView
      playCard={action('play card')}
      style={{ width: '60vw', height: '95vh' }}
      match={match}
      players={toPlayerMap(match.players, data.userMap)}
    />
  );
};
