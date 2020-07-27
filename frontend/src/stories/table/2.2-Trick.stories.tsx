import React from 'react';
import TrickView from 'features/table/TrickView';
import { Diamond10, Diamond9, DiamondA, DiamondQ, SpadeA } from 'model/cards';

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
        forehand: 'pl0',
        cards: { pl0: Diamond9, pl1: Diamond10, pl2: DiamondA, me: DiamondQ },
      }}
      players={['me', 'pl0', 'pl1', 'pl2']}
    />
  </div>
);

export const PartialTrick = () => (
  <div style={{ height: '400px' }}>
    <TrickView
      center={['50%', '50%']}
      cardWidth={150}
      trick={{ forehand: 'pl2', cards: { pl2: DiamondA, me: SpadeA } }}
      players={['me', 'pl0', 'pl1', 'pl2']}
    />
  </div>
);
