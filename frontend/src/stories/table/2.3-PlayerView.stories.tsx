import React from 'react';
import PlayerView from 'features/table/PlayerView';
import { User } from 'model/core';
import { BidType, GameType } from 'api/karlchen_pb';

/* eslint import/no-anonymous-default-export: [2, {"allowObject": true}] */
export default {
  title: 'Match/Player',
  component: PlayerView,
};

const user: User = { id: 'pl0', name: 'Woldemar', online: true };
export const Plain = () => (
  <div style={{ maxWidth: 250 }}>
    <PlayerView user={user} turn={false} bids={[]} />
  </div>
);

export const Turn = () => (
  <div style={{ maxWidth: 250 }}>
    <PlayerView user={user} turn bids={[]} />
  </div>
);

export const Offline = () => (
  <div style={{ maxWidth: 250 }}>
    <PlayerView user={{ ...user, online: false }} turn={false} bids={[]} />
  </div>
);

export const Solo = () => (
  <div style={{ maxWidth: 250 }}>
    <PlayerView user={user} turn={false} bids={[]} soloGame={GameType.HEARTS_SOLO} />
  </div>
);

export const ReKeine90 = () => (
  <div style={{ maxWidth: 250 }}>
    <PlayerView user={user} turn={false} bids={[BidType.RE_BID, BidType.RE_NO_NINETY]} />
  </div>
);
