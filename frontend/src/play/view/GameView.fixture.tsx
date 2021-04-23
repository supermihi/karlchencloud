import React from 'react';
import GameView from './GameView';
import { Match, Game } from 'model/match';
import { MatchPhase, GameType, BidType } from 'api/karlchen_pb';
import { Pos, toPlayerMap } from 'model/players';
import { emptyAuction } from 'model/auction';
import * as mockCards from 'mocks/cards';
import * as mockPlayers from 'mocks/players';
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
  currentTrick: mockCards.trick(Pos.left, 3),
};

const match: Match = {
  phase: MatchPhase.GAME,
  players: mockPlayers.players,
  cards: mockCards.fullHand,
  game: game,
  auction: emptyAuction(),
  turn: Pos.bottom,
};
export default (
  <div style={{ width: '75vw', height: '95vh' }}>
    <GameView match={match} players={toPlayerMap(match.players, mockPlayers.userMap)} />
  </div>
);
