import React from 'react';
import { Trick } from 'model/match';
import { getCardUrl } from 'shared/Cards';
import { cardString } from 'model/cards';
import { nthNext, Pos } from 'model/players';
import { makeStyles } from '@material-ui/core/styles';
import clsx from 'clsx';
import { sleep } from 'shared/sleep';

interface Props {
  trick: Trick;
  cardWidth: number;
  center: string[];
}

const useStyles = makeStyles(() => ({
  [Pos.top]: {
    animation: '$playCardTop 1000ms',
    animationFillMode: 'forwards',
  },
  '@keyframes playCardTop': {
    from: {
      top: '-20%',
    },
  },
  [Pos.right]: {
    animation: '$playCardRight 1000ms',
    animationFillMode: 'forwards',
  },
  '@keyframes playCardRight': {
    from: {
      left: '120%',
    },
  },
  [Pos.bottom]: {
    animation: '$playCardBottom 1000ms',
    animationFillMode: 'forwards',
  },
  '@keyframes playCardBottom': {
    from: {
      top: '120%',
    },
  },
  [Pos.left]: {
    animation: '$playCardLeft 1000ms',
    animationFillMode: 'forwards',
  },
  '@keyframes playCardLeft': {
    from: {
      left: '-20%',
    },
  },
  card: {},
  fadeBottom: {
    animation: `$fadeBottom 2000ms`,
    animationFillMode: 'forwards',
  },
  '@keyframes fadeBottom': {
    '100%': {
      top: '120%',
    },
  },
  fadeTop: {
    animation: `$fadeTop 2000ms`,
    animationFillMode: 'forwards',
  },
  '@keyframes fadeTop': {
    '100%': {
      top: '-20%',
    },
  },
  fadeLeft: {
    animation: `$fadeLeft 2000ms`,
    animationFillMode: 'forwards',
  },
  '@keyframes fadeLeft': {
    '100%': {
      left: '-20%',
    },
  },
  fadeRight: {
    animation: `fadeRight 2000ms`,
    animationFillMode: 'forwards',
  },
  '@keyframes fadeRight': {
    '100%': {
      left: '120%',
    },
  },
}));

export default function TrickView({
  trick: { cards, forehand, winner },
  cardWidth,
  center: [x, y],
}: Props): React.ReactElement {
  const classes = useStyles();
  const [fade, setFade] = React.useState(false);
  React.useEffect(() => {
    const func = async () => {
      if (winner != null) {
        await sleep(1000);
        setFade(true);
      } else {
        setFade(false);
      }
    };
    func();
  }, [cards, winner, fade]);
  return (
    <>
      {cards.map((card, i) => {
        const pos = nthNext(forehand, i);
        return (
          <div
            key={`${i}-${card.rank}-${card.suit}-${winner}`}
            style={{ left: x, top: y, position: 'absolute' }}
            className={clsx({
              [classes[pos]]: i === cards.length - 1,
              [classes.fadeBottom]: fade && winner === Pos.bottom,
              [classes.fadeLeft]: fade && winner === Pos.left,
              [classes.fadeTop]: fade && winner === Pos.top,
              [classes.fadeRight]: fade && winner === Pos.right,
            })}
          >
            <img
              alt={cardString(card)}
              src={getCardUrl(card)}
              width={cardWidth}
              className={classes.card}
              style={{
                transform: cardTransform(pos),
              }}
            />
          </div>
        );
      })}
    </>
  );
}

function cardTransform(pos: Pos): string {
  switch (pos) {
    case Pos.top:
      return `translate(-50%, -50%) rotate(5deg) translate(0, -25%)`;
    case Pos.right:
      return `translate(-50%, -50%) rotate(87deg) translate(0, -25%)`;
    case Pos.bottom:
      return `translate(-50%, -50%) rotate(182deg) translate(0, -25%)`;
    case Pos.left:
      return `translate(-50% ,-50%) rotate(276deg) translate(0,-25%)`;
  }
}
