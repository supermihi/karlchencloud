import React from 'react';
import { Declaration } from 'model/auction';
import { BidType, GameType } from 'api/karlchen_pb';
import * as mockPlayers from 'mocks/players';
import PlayerView from './PlayerView';

export default {
  normal: <PlayerView user={mockPlayers.face} bids={[]} turn={true} />,
  solo: (
    <PlayerView user={mockPlayers.left} soloGame={GameType.DIAMONDS_SOLO} bids={[]} turn={true} />
  ),
  declaration: (
    <PlayerView user={mockPlayers.right} declaration={Declaration.gesund} bids={[]} turn={false} />
  ),

  bids: (
    <PlayerView
      user={mockPlayers.right}
      bids={[BidType.RE_BID, BidType.RE_NO_NINETY, BidType.RE_NO_SIXTY]}
      turn={false}
    />
  ),
};
