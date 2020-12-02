import React from 'react';
import { Rank, Suit } from 'api/karlchen_pb';
import { action } from '@storybook/addon-actions';
import OwnCardsView from 'features/table/OwnCardsView';

/* eslint import/no-anonymous-default-export: [2, {"allowObject": true}] */
export default {
  title: 'Match/Cards',
  component: OwnCardsView,
};

export const OwnCards = () => (
  <div style={{ width: '50%', minHeight: '180px', overflowY: 'hidden', position: 'relative' }}>
    <OwnCardsView
      cardWidth={120}
      cards={[
        { rank: Rank.ACE, suit: Suit.DIAMONDS },
        { rank: Rank.KING, suit: Suit.CLUBS },
        { rank: Rank.QUEEN, suit: Suit.HEARTS },
        { rank: Rank.QUEEN, suit: Suit.HEARTS },
        { rank: Rank.QUEEN, suit: Suit.HEARTS },
        { rank: Rank.QUEEN, suit: Suit.HEARTS },
        { rank: Rank.QUEEN, suit: Suit.HEARTS },
      ]}
      onClick={action('clicked card')}
    />
  </div>
);
