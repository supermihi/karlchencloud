import React from 'react';
import PlayerView from 'features/table/PlayerView';

export default {
  title: 'Match/Player',
  component: PlayerView,
};

export const Player = () => (
  <div style={{ maxWidth: 250 }}>
    <PlayerView user={{ id: 'pl0', name: 'Woldemar', online: true }} />
  </div>
);
