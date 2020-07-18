import React from 'react';
import { Trick } from 'model/match';
import {Players} from "model/table";
import {Card} from "../../model/core";
import SvgCard, {svgCardHeight, svgCardWidth} from "../../components/SvgCards";

interface Props {
  trick: Trick;
  players: Players;
  cardHeight: string | number;
}

export default function TrickView({ trick, players, cardHeight }: Props) {
  const forehandIndex = players.indexOf(trick.forehand);
  const forehandRotation = (180+90*forehandIndex);
  const excenter = svgCardWidth / 5;
  const horizontalPlus = (svgCardHeight + 2*excenter) - svgCardWidth;
  const size = svgCardHeight + 2*excenter;
  const height = `calc(${(size / svgCardHeight)} * ${cardHeight})`;
  const cards: Card[] = [];
  for (let i = 0; i < players.length; i += 1) {
    const player = players[(forehandIndex + i) % 4];
    if (!trick.cards[player])
      break;
    cards.push(trick.cards[player]);
  }
  return <svg height={height} viewBox={`${-horizontalPlus/2} ${-excenter} ${size} ${size}`}>
    {cards.map((card, i) => <SvgCard card={card}  transform={`translate(0,${-excenter}) rotate(${(forehandRotation + 90*i) % 360},${svgCardWidth/2},${svgCardHeight/2+excenter})`}/>)}

  </svg>
}
