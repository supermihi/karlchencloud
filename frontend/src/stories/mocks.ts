import { User, toUserMap } from 'model/core';
import * as cards from 'model/cards';
import { Players } from 'model/table';
import { Trick } from 'model/match';
import { zipObject } from 'lodash';

export const me: User = { id: 'me', name: 'Ich', online: true };
export const left: User = { id: 'left', name: 'Spieler Links', online: false };
export const face: User = { id: 'face', name: 'Spieler Gegenüber', online: true };
export const right: User = { id: 'right', name: 'Spieler Rechts', online: true };

export const users = [right, me, face, left];
export const userMap = toUserMap(users);

export const players: Players = [me.id, left.id, face.id, right.id];
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

export function trick(forehandIndex: number, trickSize: number): Trick {
  return {
    forehand: players[forehandIndex],
    cards: zipObject(
      [...players, ...players].slice(forehandIndex, forehandIndex + trickSize),
      trickCards.slice(0, trickSize)
    ),
  };
}
