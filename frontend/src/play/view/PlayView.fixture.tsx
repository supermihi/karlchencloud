import React from 'react';
import PlayView from 'play/view/PlayView';
import { matchInAuction, matchInGame, table } from './mocks';
import { useSelect } from 'react-cosmos/fixture';

function PlayViewFixture(): React.ReactElement {
  const [phase] = useSelect('phase', {
    options: ['auction', 'game'],
  });
  return (
    <div style={{ width: '75vw', height: '95vh' }}>
      <PlayView
        match={phase === 'auction' ? matchInAuction : matchInGame}
        table={table}
        dispatch={(action: unknown) => console.log(action)}
      />
    </div>
  );
}
export default PlayViewFixture;
