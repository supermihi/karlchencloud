import Diamonds9 from './resources/cards/9D.svg';
import Heart9 from './resources/cards/9H.svg';
import Spade9 from './resources/cards/9S.svg';
import Club9 from './resources/cards/9C.svg';

import Diamonds10 from './resources/cards/TD.svg';
import Heart10 from './resources/cards/TH.svg';
import Spade10 from './resources/cards/TS.svg';
import Club10 from './resources/cards/TC.svg';

import DiamondsJ from './resources/cards/JD.svg';
import HeartJ from './resources/cards/JH.svg';
import SpadeJ from './resources/cards/JS.svg';
import ClubJ from './resources/cards/JC.svg';

import DiamondsQ from './resources/cards/QD.svg';
import HeartQ from './resources/cards/QH.svg';
import SpadeQ from './resources/cards/QS.svg';
import ClubQ from './resources/cards/QC.svg';

import DiamondsK from './resources/cards/KD.svg';
import HeartK from './resources/cards/KH.svg';
import SpadeK from './resources/cards/KS.svg';
import ClubK from './resources/cards/KC.svg';

import DiamondsA from './resources/cards/AD.svg';
import HeartA from './resources/cards/AH.svg';
import SpadeA from './resources/cards/AS.svg';
import ClubA from './resources/cards/AC.svg';
import { Rank, Suit } from 'api/karlchen_pb';
import { Card } from 'model/core';

export function getCardUrl(card: Card) {
  switch (card.rank) {
    case Rank.NINE:
      switch (card.suit) {
        case Suit.DIAMONDS:
          return Diamonds9;
        case Suit.HEARTS:
          return Heart9;
        case Suit.SPADES:
          return Spade9;
        case Suit.CLUBS:
          return Club9;
      }
      break;
    case Rank.TEN:
      switch (card.suit) {
        case Suit.DIAMONDS:
          return Diamonds10;
        case Suit.HEARTS:
          return Heart10;
        case Suit.SPADES:
          return Spade10;
        case Suit.CLUBS:
          return Club10;
      }
      break;
    case Rank.JACK:
      switch (card.suit) {
        case Suit.DIAMONDS:
          return DiamondsJ;
        case Suit.HEARTS:
          return HeartJ;
        case Suit.SPADES:
          return SpadeJ;
        case Suit.CLUBS:
          return ClubJ;
      }
      break;
    case Rank.QUEEN:
      switch (card.suit) {
        case Suit.DIAMONDS:
          return DiamondsQ;
        case Suit.HEARTS:
          return HeartQ;
        case Suit.SPADES:
          return SpadeQ;
        case Suit.CLUBS:
          return ClubQ;
      }
      break;
    case Rank.KING:
      switch (card.suit) {
        case Suit.DIAMONDS:
          return DiamondsK;
        case Suit.HEARTS:
          return HeartK;
        case Suit.SPADES:
          return SpadeK;
        case Suit.CLUBS:
          return ClubK;
      }
      break;
    case Rank.ACE:
      switch (card.suit) {
        case Suit.DIAMONDS:
          return DiamondsA;
        case Suit.HEARTS:
          return HeartA;
        case Suit.SPADES:
          return SpadeA;
        case Suit.CLUBS:
          return ClubA;
      }
  }
}
