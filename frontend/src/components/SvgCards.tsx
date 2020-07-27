import React, { SVGProps } from 'react';
import { ReactComponent as Diamonds9 } from './resources/cards/9D.svg';
import { ReactComponent as Heart9 } from './resources/cards/9H.svg';
import { ReactComponent as Spade9 } from './resources/cards/9S.svg';
import { ReactComponent as Club9 } from './resources/cards/9C.svg';

import { ReactComponent as Diamonds10 } from './resources/cards/TD.svg';
import { ReactComponent as Heart10 } from './resources/cards/TH.svg';
import { ReactComponent as Spade10 } from './resources/cards/TS.svg';
import { ReactComponent as Club10 } from './resources/cards/TC.svg';

import { ReactComponent as DiamondsJ } from './resources/cards/JD.svg';
import { ReactComponent as HeartJ } from './resources/cards/JH.svg';
import { ReactComponent as SpadeJ } from './resources/cards/JS.svg';
import { ReactComponent as ClubJ } from './resources/cards/JC.svg';

import { ReactComponent as DiamondsQ } from './resources/cards/QD.svg';
import { ReactComponent as HeartQ } from './resources/cards/QH.svg';
import { ReactComponent as SpadeQ } from './resources/cards/QS.svg';
import { ReactComponent as ClubQ } from './resources/cards/QC.svg';

import { ReactComponent as DiamondsK } from './resources/cards/KD.svg';
import { ReactComponent as HeartK } from './resources/cards/KH.svg';
import { ReactComponent as SpadeK } from './resources/cards/KS.svg';
import { ReactComponent as ClubK } from './resources/cards/KC.svg';

import { ReactComponent as DiamondsA } from './resources/cards/AD.svg';
import { ReactComponent as HeartA } from './resources/cards/AH.svg';
import { ReactComponent as SpadeA } from './resources/cards/AS.svg';
import { ReactComponent as ClubA } from './resources/cards/AC.svg';

import { Card } from '../model/core';
import { Rank, Suit } from '../api/karlchen_pb';

export const svgCardWidth = 212;
export const svgCardHeight = 329;
export const cardAspectRatio = svgCardWidth / svgCardHeight;

interface Props extends SVGProps<SVGSVGElement> {
  card: Card;
}
export default function SvgCard({ card, ...props }: Props) {
  const CardComponent = getComponent(card);
  return <CardComponent {...props} />;
  /*return <svg {...props}>
    <rect width={svgCardWidth} height={svgCardHeight}
          style={{stroke: 'black', fill: 'white', strokeWidth: 3}}
          rx={35}
          ry={35}
    />
    <CardComponent />
  </svg>;*/
}
function getComponent(card: Card) {
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
      break;
  }
  throw new Error('unexpected rank');
}
