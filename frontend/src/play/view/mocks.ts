import { BidType, GameType, MatchPhase, TablePhase } from '../../api/karlchen_pb';
import { Pos } from 'model/players';
import { Game, Match } from 'model/match';
import * as mockCards from 'mocks/cards';
import { Table } from 'model/table';
import * as mockPlayers from 'mocks/players';
import { Auction, Declaration, emptyAuction } from 'model/auction';

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

const auction: Auction = {
  declarations: { [Pos.top]: Declaration.gesund, [Pos.right]: Declaration.vorbehalt },
};
export const matchInGame: Match = {
  phase: MatchPhase.GAME,
  players: mockPlayers.players,
  cards: mockCards.fullHand,
  game: game,
  auction: emptyAuction(),
  turn: Pos.bottom,
};
export const matchInAuction: Match = {
  phase: MatchPhase.AUCTION,
  players: mockPlayers.players,
  cards: mockCards.fullHand,
  game: null,
  auction,
  turn: Pos.bottom,
};

export const table: Table = {
  id: '123',
  members: mockPlayers.users,
  owner: mockPlayers.me.id,
  created: '1',
  phase: TablePhase.PLAYING,
};
