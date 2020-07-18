import React from 'react';
import { Card } from 'model/core';
import SvgCard, {cardAspectRatio, svgCardHeight, svgCardWidth} from "components/SvgCards";
interface Props {
  cards: Card[];
  cardHeight: string | number;
  onClick?: (card: Card, index: number) => void;
}
export default function ({ cards, onClick, cardHeight }: Props) {
  const translate = svgCardWidth / 5;
  const width = svgCardWidth + (cards.length - 1) * translate;
  console.log(`width ${width} svgHeight ${svgCardHeight} svgWidth ${svgCardWidth} AR ${cardAspectRatio}`)
  return (
      <svg viewBox={`0 0 ${width} ${svgCardHeight}`} height={cardHeight}>
        {cards.map((card, i) => (
          <SvgCard style={{cursor: "grab"}}
                   onClick={_ => onClick!== undefined && onClick(card, i)}
                   card={card}
                   key={`card${i}`}
                   width={svgCardWidth}
                   height={svgCardHeight}
                   transform={`translate(${translate * i},0)`} />
        ))}
      </svg>
  );
}
