import * as api from 'api/karlchen_pb';
import { Table } from './table';
import { User, Card } from './core';
import { fromPairs, mapValues, groupBy } from 'lodash';
import { Game, Match, Mode, PlayedCard, Trick } from './match';
import { toDate } from 'api/helpers';
import { getPosition, Pos, PlayerIds } from './players';
import { Auction, Declaration, DeclareResult, emptyAuction } from './auction';

export interface TableState {
  table: Table;
  match: Match | null;
}
export function toTable(t: api.TableData, phase: api.TablePhase): Table {
  return {
    owner: t.getOwner(),
    id: t.getTableId(),
    invite: t.getInviteCode(),
    members: t.getMembersList().map(toUser),
    created: toDate(t.getCreated()!).toLocaleString(),
    phase,
  };
}

export function getCurrentTableState(u: api.UserState): TableState | null {
  return u.hasCurrenttable() ? toTableState(u.getCurrenttable()!) : null;
}

export function toTableState(t: api.TableState): TableState {
  return {
    table: toTable(t.getData()!, t.getPhase()),
    match: t.hasCurrentMatch() ? toMatch(t.getCurrentMatch()!) : null,
  };
}

export function toUser(member: api.TableMember): User {
  return {
    id: member.getUserId(),
    name: member.getName(),
    online: member.getOnline(),
  };
}

function toPlayers(p: api.Players): PlayerIds {
  return {
    [Pos.bottom]: p.getUserIdSelf(),
    [Pos.left]: p.getUserIdLeft(),
    [Pos.top]: p.getUserIdFace(),
    [Pos.right]: p.getUserIdRight(),
  };
}
export function toAuction(players: PlayerIds, m: api.AuctionState): Auction {
  return {
    declarations: fromPairs(
      m
        .getDeclarationsList()
        .map((d) => [
          getPosition(players, d.getUserId()),
          d.getVorbehalt() ? Declaration.vorbehalt : Declaration.gesund,
        ])
    ),
  };
}

function getPositionOptional(players: PlayerIds, id?: string) {
  if (id === undefined) return undefined;
  return getPosition(players, id);
}
export function toMode(m: api.Mode, players: PlayerIds): Mode {
  return {
    type: m.getType(),
    soloist: getPositionOptional(players, m.getSoloist()?.getUserId()),
    spouse: getPositionOptional(players, m.getSpouse()?.getUserId()),
    forehand: getPosition(players, m.getForehand()),
  };
}
function toGame(g: api.GameState, players: PlayerIds): Game {
  const bidsById = mapValues(
    groupBy(g.getBidsList(), (b) => b.getUserId()),
    (bids) => bids.map((bid) => bid.getBid())
  );
  const bids = mapValues(players, (id) => bidsById[id] || []);
  return {
    bids,
    completedTricks: g.getCompletedTricks(),
    currentTrick: toTrick(g.getCurrentTrick() as api.Trick, players),
    previousTrick: g.hasPreviousTrick()
      ? toTrick(g.getPreviousTrick() as api.Trick, players)
      : undefined,
    mode: toMode(g.getMode() as api.Mode, players),
  };
}

function toTrick(t: api.Trick, players: PlayerIds): Trick {
  const forehand = getPosition(players, t.getUserIdForehand());
  let winner: Pos | null = null;
  if (t.hasUserIdWinner()) {
    winner = getPosition(players, t.getUserIdWinner()!.getUserId());
  }
  const cards = t.getCardsList().map(toCard);
  return { forehand, cards, winner };
}
export function toCard(c: api.Card): Card {
  return {
    suit: c.getSuit(),
    rank: c.getRank(),
  };
}

export function toMatch(m: api.MatchState): Match {
  const players = toPlayers(m.getPlayers()!);
  const cards = m.hasOwnCards() ? m.getOwnCards()!.getCardsList().map(toCard) : [];
  const auction = m.hasAuctionState() ? toAuction(players, m.getAuctionState()!) : emptyAuction();
  const game = m.hasGameState() ? toGame(m.getGameState()!, players) : null;
  return {
    phase: m.getPhase(),
    turn: m.hasTurn() ? getPosition(players, m.getTurn()!.getUserId()) : null,
    players,
    cards,
    auction,
    game,
    winner: null,
  };
}

export function toDeclareResult(decl: api.Declaration, players: PlayerIds): DeclareResult {
  const apiMode = decl.getDefinedgamemode();
  const mode = apiMode === undefined ? null : toMode(apiMode, players);
  return {
    mode,
    player: getPosition(players, decl.getUserId()),
    declaration: decl.getVorbehalt() ? Declaration.vorbehalt : Declaration.gesund,
  };
}

export function toPlayedCard(card: api.PlayedCard, players: PlayerIds): PlayedCard {
  return {
    card: toCard(card.getCard()!),
    player: getPosition(players, card.getUserId()),
    trickWinner: card.hasTrickWinner()
      ? getPosition(players, card.getTrickWinner()!.getUserId())
      : null,
    matchWinner: card.hasWinner() ? card.getWinner()!.getParty() : null,
  };
}
