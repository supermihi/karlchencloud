import * as api from 'api/karlchen_pb';
export interface User {
  name: string;
  id: string;
  online?: boolean;
}

export interface Card {
  suit: api.Suit;
  rank: api.Rank;
}

export function cardString({ suit, rank }: Card): string {
  return `${suitString(suit)}${rankString(rank)}`;
}

function rankString(rank: api.Rank): string {
  switch (rank) {
    case api.Rank.ACE:
      return 'A';
    case api.Rank.JACK:
      return 'J';
    case api.Rank.KING:
      return 'K';
    case api.Rank.QUEEN:
      return 'Q';
    case api.Rank.NINE:
      return '9';
    case api.Rank.TEN:
      return '10';
  }
}

function suitString(suit: api.Suit): string {
  switch (suit) {
    case api.Suit.DIAMONDS:
      return '♦';
    case api.Suit.HEARTS:
      return '♥';
    case api.Suit.SPADES:
      return '♠';
    case api.Suit.CLUBS:
      return '♣';
  }
}
