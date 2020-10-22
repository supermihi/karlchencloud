import { User, toUserMap } from 'model/core';
import * as cards from 'model/cards';
import { Trick } from 'model/match';
import { PlayerIds, Pos } from 'model/players';

export const me: User = { id: 'me', name: 'Ich', online: true };
export const left: User = { id: 'left', name: 'Spieler Links', online: false };
export const face: User = { id: 'face', name: 'Spieler Gegenüber', online: true };
export const right: User = { id: 'right', name: 'Spieler Rechts', online: true };

export const users = [right, me, face, left];
export const userMap = toUserMap(users);

export const players: PlayerIds = {
  [Pos.bottom]: me.id,
  [Pos.left]: left.id,
  [Pos.top]: face.id,
  [Pos.right]: right.id,
};
const trickCards = [cards.Club9, cards.ClubA, cards.Club10, cards.DiamondA];

export const fullHand = [
  cards.Heart9,
  cards.HeartK,
  cards.HeartA,
  cards.Spade9,
  cards.Spade10,
  cards.SpadeA,
  cards.Diamond9,
  cards.DiamondA,
  cards.ClubJ,
  cards.DiamondQ,
  cards.SpadeQ,
];

export function trick(forehand: Pos, trickSize: number): Trick {
  return { forehand, cards: trickCards.slice(0, trickSize) };
}
