import React from 'react';
import { Declaration } from 'model/auction';
import { BidType, GameType } from 'api/karlchen_pb';
import { User } from 'model/core';
import * as mockPlayers from 'mocks/players';
import PlayerView from './PlayerView';


interface Props {
  user: User;
  declaration?: Declaration;
  bids: BidType[];
  soloGame?: GameType;
  turn: boolean;
}

const Player = ({ user, declaration, bids, soloGame, turn }: Props) => (
  <div style={{ height: '400px' }}>
    <PlayerView
      user={user}
      soloGame={soloGame}
      declaration={declaration}
      bids={bids}
      turn={turn}
    />
  </div>
);

export default {
  'normal': <Player user={mockPlayers.face}
                        bids={[]} turn={true} />,
  'solo': <Player user={mockPlayers.left} soloGame={GameType.DIAMONDS_SOLO}
                      bids={[]} turn={true} />,
  'declaration': <Player user={mockPlayers.right} declaration={Declaration.gesund}
                             bids={[]} turn={false} />,
};
