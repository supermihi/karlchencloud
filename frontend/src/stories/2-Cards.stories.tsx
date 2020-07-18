import React from 'react';
import CardsView from 'features/table/CardsView';
import { Rank, Suit } from 'api/karlchen_pb';
import {action} from "@storybook/addon-actions";

export default {
  title: 'Cards',
  component: CardsView,
};
export const Card = () => (
  <CardsView
    cards={[
      { rank: Rank.ACE, suit: Suit.DIAMONDS },
      { rank: Rank.KING, suit: Suit.CLUBS },
      { rank: Rank.QUEEN, suit: Suit.HEARTS},
    ]}
    onClick={action('click card')}
  />
);
