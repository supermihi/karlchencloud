import React from 'react';
import MatchView from 'features/table/MatchView';
import { Match, Game } from 'model/match';
import { MatchPhase, GameType, BidType } from 'api/karlchen_pb';
import * as data from './mocks';
import { Pos, toPlayerMap } from 'model/players';
import { action } from '@storybook/addon-actions';
import { emptyAuction } from 'model/auction';
import { withProvider } from '../provider';

/* eslint import/no-anonymous-default-export: [2, {"allowObject": true}] */
export default {
  title: 'Match/Match',
  component: MatchView,
  decorators: [withProvider],
};

const mode = {
  type: GameType.NORMAL_GAME,
  forehand: Pos.bottom,
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
  game: game,
  auction: emptyAuction(),
  turn: Pos.bottom,
};
export const MatchTable = () => {
  return (
    <div style={{ width: '60vw', height: '95vh' }}>
      <MatchView
        playCard={action('play card')}
        match={match}
        players={toPlayerMap(match.players, data.userMap)}
      />
    </div>
  );
};
