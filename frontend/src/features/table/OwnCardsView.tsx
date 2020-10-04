import React, { useState } from 'react';
import { Card } from 'model/core';
import { getCardUrl } from 'components/UrlCards';
import { cardAspectRatio } from 'components/SvgCards';
import { cardString } from 'model/cards';

interface Props {
  cards: Card[];
  cardWidth: number;
  onClick?: (card: Card, index: number) => void;
}

export default function OwnCardsView({ cards, cardWidth, onClick }: Props) {
  const angle = 2.5;
  const [hoveredIndex, setHoveredIndex] = useState(-1);
  const allowClick = Boolean(onClick);
  return (
    <>
      {cards.map((card, i) => (
        <img
          key={`${card.rank}/${card.suit}/${i}`}
          alt={cardString(card)}
          src={getCardUrl(card)}
          width={cardWidth}
          onMouseEnter={(_) => setHoveredIndex(i)}
          onMouseLeave={(_) => setHoveredIndex(-1)}
          height={cardWidth / cardAspectRatio}
          style={{
            position: 'absolute',
            transition: 'transform .2s',
            bottom: (-0.5 * cardWidth) / cardAspectRatio,
            left: `calc(50% - ${cardWidth / 2}px)`,
            transform: `rotate(${angle * (i + 1 - cards.length / 2)}deg) ${
              allowClick && hoveredIndex === i ? 'translate(0, -20px)' : ''
            }`,
            transformOrigin: `${cardWidth / 2}px ${(cardWidth / cardAspectRatio) * 4}px`,
            cursor: allowClick ? 'pointer' : 'inherit',
          }}
          onClick={onClick && ((_) => onClick(card, i))}
        />
      ))}
    </>
  );
}
