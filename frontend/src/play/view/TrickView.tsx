import React from 'react';
import { Trick } from 'model/match';
import { getCardUrl } from 'shared/Cards';
import { cardString } from 'model/cards';
import { nthNext, Pos } from 'model/players';

interface Props {
  trick: Trick;
  cardWidth: number;
  center: string[];
}

function fadeTarget(defaultValue: string, horizontal: boolean, winner: Pos) {
  if (
    (horizontal && (winner === Pos.top || winner === Pos.bottom)) ||
    (!horizontal && (winner === Pos.left || winner === Pos.right))
  ) {
    return defaultValue;
  }
  if (winner === Pos.top || winner === Pos.left) {
    return '-50%';
  }
  return '150%';
}

export default function TrickView({
  trick: { cards, forehand, winner },
  cardWidth,
  center: [x, y],
}: Props): React.ReactElement {
  const excenter = cardWidth / 4;
  const [startTransition, setFadeOut] = React.useState(false);
  React.useEffect(() => {
    if (winner != null) {
      setFadeOut(true);
    } else {
      setFadeOut(false);
    }
  }, [winner]);
  const setFade = startTransition && winner !== null;
  return (
    <>
      {cards.map((card, i) => (
        <img
          key={`${i}-${card.rank}-${card.suit}-${winner}`}
          alt={cardString(card)}
          src={getCardUrl(card)}
          width={cardWidth}
          style={{
            left: setFade ? fadeTarget(x, true, winner!) : x,
            top: setFade ? fadeTarget(y, false, winner!) : y,
            position: 'absolute',
            transform: cardTransform(nthNext(forehand, i), excenter),
            transition: winner !== null ? 'left 1s, top 1s' : '',
            transitionDelay: '0.7s',
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
