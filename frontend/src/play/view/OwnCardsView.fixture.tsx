import React from 'react';
import { Rank, Suit } from 'api/karlchen_pb';
import OwnCardsView from './OwnCardsView';
interface Props {
  interactive: boolean;
}
const OwnCards = ({ interactive }: Props) => (
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
      onClick={interactive ? () => 0 : undefined}
    />
  </div>
);
export default <OwnCards interactive />;
