import * as cards from 'model/cards';
import { Trick } from 'model/match';
import { Pos } from 'model/players';

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

export function trick(forehand: Pos, trickSize: number, trickWinner?: Pos): Trick {
  return {
    forehand,
    cards: trickCards.slice(0, trickSize),
    winner: trickWinner ?? null,
  };
}
