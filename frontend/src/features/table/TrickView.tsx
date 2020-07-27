import React from 'react';
import { Trick } from 'model/match';
import { Players } from 'model/table';
import { Card } from '../../model/core';
import { getCardUrl } from 'components/UrlCards';
import { cardString } from 'model/cards';

interface Props {
  trick: Trick;
  players: Players;
  cardWidth: number;
  center: [string, string];
}

export default function TrickView({ trick, players, cardWidth, center: [x, y] }: Props) {
  const forehandIndex = players.indexOf(trick.forehand);
  const excenter = cardWidth / 4;
  const cards: Card[] = [];
  for (let i = 0; i < players.length; i += 1) {
    const player = players[(forehandIndex + i) % 4];
    if (!trick.cards[player]) break;
    cards.push(trick.cards[player]);
  }

  return (
    <>
      {cards.map((card, i) => (
        <img
          alt={cardString(card)}
          src={getCardUrl(card)}
          width={cardWidth}
          style={{
            left: x,
            top: y,
            position: 'absolute',
            transform: cardTransform((forehandIndex + i) % 4, excenter),
          }}
        />
      ))}
    </>
  );
}

function cardTransform(i: number, excenter: number): string {
  switch (i) {
    case 2:
      return `translate(-50%, calc(-50% - ${excenter}px)) rotate(5deg)`;
    case 3:
      return `translate(calc(-50% + ${excenter}px), -50%) rotate(87deg)`;
    case 0:
      return `translate(-50%, calc(-50% + ${excenter}px)) rotate(182deg)`;
    case 1:
      return `translate(calc(-50% - ${excenter}px), -50%) rotate(276deg)`;
  }
  throw new Error(`invalid i in cardTranstorm: ${i}`);
}
