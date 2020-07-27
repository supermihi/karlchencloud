import { Card } from './core';
import { Rank, Suit } from 'api/karlchen_pb';

export const Diamond9: Card = { rank: Rank.NINE, suit: Suit.DIAMONDS };
export const Diamond10: Card = { rank: Rank.TEN, suit: Suit.DIAMONDS };
export const DiamondJ: Card = { rank: Rank.JACK, suit: Suit.DIAMONDS };
export const DiamondQ: Card = { rank: Rank.QUEEN, suit: Suit.DIAMONDS };
export const DiamondK: Card = { rank: Rank.KING, suit: Suit.DIAMONDS };
export const DiamondA: Card = { rank: Rank.ACE, suit: Suit.DIAMONDS };

export const Heart9: Card = { rank: Rank.NINE, suit: Suit.HEARTS };
export const Heart10: Card = { rank: Rank.TEN, suit: Suit.HEARTS };
export const HeartJ: Card = { rank: Rank.JACK, suit: Suit.HEARTS };
export const HeartQ: Card = { rank: Rank.QUEEN, suit: Suit.HEARTS };
export const HeartK: Card = { rank: Rank.KING, suit: Suit.HEARTS };
export const HeartA: Card = { rank: Rank.ACE, suit: Suit.HEARTS };

export const Spade9: Card = { rank: Rank.NINE, suit: Suit.SPADES };
export const Spade10: Card = { rank: Rank.TEN, suit: Suit.SPADES };
export const SpadeJ: Card = { rank: Rank.JACK, suit: Suit.SPADES };
export const SpadeQ: Card = { rank: Rank.QUEEN, suit: Suit.SPADES };
export const SpadeK: Card = { rank: Rank.KING, suit: Suit.SPADES };
export const SpadeA: Card = { rank: Rank.ACE, suit: Suit.SPADES };

export const Club9: Card = { rank: Rank.NINE, suit: Suit.CLUBS };
export const Club10: Card = { rank: Rank.TEN, suit: Suit.CLUBS };
export const ClubJ: Card = { rank: Rank.JACK, suit: Suit.CLUBS };
export const ClubQ: Card = { rank: Rank.QUEEN, suit: Suit.CLUBS };
export const ClubK: Card = { rank: Rank.KING, suit: Suit.CLUBS };
export const ClubA: Card = { rank: Rank.ACE, suit: Suit.CLUBS };

export function cardString({ suit, rank }: Card): string {
  return `${suitString(suit)}${rankString(rank)}`;
}

function rankString(rank: Rank): string {
  switch (rank) {
    case Rank.ACE:
      return 'A';
    case Rank.JACK:
      return 'J';
    case Rank.KING:
      return 'K';
    case Rank.QUEEN:
      return 'Q';
    case Rank.NINE:
      return '9';
    case Rank.TEN:
      return '10';
  }
}

function suitString(suit: Suit): string {
  switch (suit) {
    case Suit.DIAMONDS:
      return '♦';
    case Suit.HEARTS:
      return '♥';
    case Suit.SPADES:
      return '♠';
    case Suit.CLUBS:
      return '♣';
  }
}
