import * as api from 'api/karlchen_pb';
import { Table, TableState } from './table';
import { User, Card } from './core';
import { fromPairs, mapValues, groupBy } from 'lodash';
import { Auction, Declaration, Game, Match, Mode, Trick } from './match';
import { toDate } from 'api/helpers';
import { getPosition, Pos, Players } from './players';

export function toTable(t: api.TableData): Table {
  return {
    owner: t.getOwner(),
    id: t.getTableId(),
    invite: t.getInviteCode(),
    players: t.getMembersList().map(toUser),
    created: toDate(t.getCreated() as api.Timestamp).toLocaleString(),
  };
}

export function getCurrentTableState(u: api.UserState): TableState | null {
  return u.hasCurrenttable() ? toTableState(u.getCurrenttable() as api.TableState) : null;
}

export function toTableState(t: api.TableState): TableState {
  return {
    table: toTable(t.getData() as api.TableData),
    match: t.hasCurrentMatch() ? toMatch(t.getCurrentMatch() as api.MatchState) : undefined,
    phase: t.getPhase(),
  };
}

export function toUser(member: api.TableMember): User {
  return {
    id: member.getUserId(),
    name: member.getName(),
    online: member.getOnline(),
  };
}

function toPlayers(p: api.Players): Players {
  return {
    [Pos.bottom]: p.getUserIdSelf(),
    [Pos.left]: p.getUserIdLeft(),
    [Pos.top]: p.getUserIdFace(),
    [Pos.right]: p.getUserIdRight(),
  };
}
export function toAuction(m: api.AuctionState): Auction {
  return {
    declarations: fromPairs(
      m
        .getDeclarationsList()
        .map((d) => [d.getUserId(), d.getVorbehalt() ? Declaration.vorbehalt : Declaration.gesund])
    ),
  };
}

function toMode(m: api.Mode): Mode {
  return {
    type: m.getType(),
    soloist: m.getSoloist()?.getUserId() ?? undefined,
    spouse: m.getSpouse()?.getUserId() ?? undefined,
    forehand: m.getForehand(),
  };
}
function toGame(g: api.GameState, players: Players): Game {
  const bidsById = mapValues(
    groupBy(g.getBidsList(), (b) => b.getUserId()),
    (bids) => bids.map((bid) => bid.getBid())
  );
  const bids = mapValues(players, (id) => bidsById[id]);
  return {
    bids,
    completedTricks: g.getCompletedTricks(),
    currentTrick: toTrick(g.getCurrentTrick() as api.Trick, players),
    previousTrick: g.hasPreviousTrick()
      ? toTrick(g.getPreviousTrick() as api.Trick, players)
      : undefined,
    mode: toMode(g.getMode() as api.Mode),
  };
}

function toTrick(t: api.Trick, players: Players): Trick {
  const forehand = getPosition(players, t.getUserIdForehand());
  let winner: Pos | undefined = undefined;
  if (t.hasUserIdWinner()) {
    winner = getPosition(players, (t.getUserIdWinner() as api.PlayerValue).getUserId());
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
  const players = toPlayers(m.getPlayers() as api.Players);
  const cards = m.hasOwnCards() ? (m.getOwnCards() as api.Cards).getCardsList().map(toCard) : [];
  return {
    phase: m.getPhase(),
    turn: m.hasTurn()
      ? getPosition(players, (m.getTurn() as api.PlayerValue).getUserId())
      : undefined,
    players,
    cards,
    details: m.hasAuctionState()
      ? toAuction(m.getAuctionState() as api.AuctionState)
      : toGame(m.getGameState() as api.GameState, players),
  };
}
