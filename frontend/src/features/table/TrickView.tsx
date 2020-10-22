import React from 'react';
import { Trick } from 'model/match';
import { getCardUrl } from 'components/UrlCards';
import { cardString } from 'model/cards';
import { nthNext, Pos } from 'model/players';

interface Props {
  trick: Trick;
  cardWidth: number;
  center: string[];
}

export default function TrickView({
  trick: { cards, forehand },
  cardWidth,
  center: [x, y],
}: Props) {
  const excenter = cardWidth / 4;
  return (
    <>
      {cards.map((card, i) => (
        <img
          key={`${i}-${card.rank}-${card.suit}`}
          alt={cardString(card)}
          src={getCardUrl(card)}
          width={cardWidth}
          style={{
            left: x,
            top: y,
            position: 'absolute',
            transform: cardTransform(nthNext(forehand, i), excenter),
          }}
        />
      ))}
    </>
  );
}

function cardTransform(pos: Pos, excenter: number): string {
  switch (pos) {
    case Pos.top:
      return `translate(-50%, calc(-50% - ${excenter}px)) rotate(5deg)`;
    case Pos.right:
      return `translate(calc(-50% + ${excenter}px), -50%) rotate(87deg)`;
    case Pos.bottom:
      return `translate(-50%, calc(-50% + ${excenter}px)) rotate(182deg)`;
    case Pos.left:
      return `translate(calc(-50% - ${excenter}px), -50%) rotate(276deg)`;
  }
}
