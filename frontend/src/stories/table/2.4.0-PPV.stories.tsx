import React from 'react';
import { Match, Game } from 'model/match';
import { MatchPhase, GameType, BidType } from 'api/karlchen_pb';
import * as data from './mocks';
import { Pos } from 'model/players';
import { emptyAuction } from 'model/auction';
import PositionedPlayerView from 'features/table/PositionedPlayerView';

/* eslint import/no-anonymous-default-export: [2, {"allowObject": true}] */
export default {
  title: 'Match/PlayerView',
  component: PositionedPlayerView,
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

export const PositionedPV = () => {
  return <PositionedPlayerView user={data.users[0]} pos={Pos.bottom} match={match} />;
};
