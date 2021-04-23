import React from 'react';
import PlayView from 'play/view/PlayView';
import { Game, Match } from 'model/match';
import { BidType, GameType, MatchPhase, TablePhase } from 'api/karlchen_pb';
import { Pos } from 'model/players';
import { emptyAuction } from 'model/auction';
import * as mockCards from 'mocks/cards';
import * as mockPlayers from 'mocks/players';
import { Table } from 'model/table';

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

const table: Table = {
  id: '123',
  members: mockPlayers.users,
  owner: mockPlayers.me.id,
  created: '1',
  phase: TablePhase.PLAYING,
};
export default (
  <div style={{ width: '75vw', height: '95vh' }}>
    <PlayView match={match} table={table} />
  </div>
);
