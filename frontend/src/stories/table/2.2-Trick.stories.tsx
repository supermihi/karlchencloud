import React from 'react';
import TrickView from 'features/table/TrickView';
import { Diamond10, Diamond9, DiamondA, DiamondQ, SpadeA } from 'model/cards';
import { Pos } from 'model/players';

export default {
  title: 'Match/Trick',
  component: TrickView,
};

export const CompleteTrick = () => (
  <div style={{ height: '400px', position: 'relative' }}>
    <TrickView
      center={['50%', '50%']}
      cardWidth={150}
      trick={{
        forehand: Pos.left,
        cards: [Diamond9, Diamond10, DiamondA, DiamondQ],
      }}
    />
  </div>
);

export const PartialTrick = () => (
  <div style={{ height: '400px' }}>
    <TrickView
      center={['50%', '50%']}
      cardWidth={150}
      trick={{ forehand: Pos.right, cards: [DiamondA, SpadeA] }}
    />
  </div>
);
