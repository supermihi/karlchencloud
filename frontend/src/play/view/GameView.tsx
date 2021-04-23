import React, { useEffect, useRef, useState } from 'react';
import { Match } from 'model/match';
import { PlayingUsers, Pos } from 'model/players';
import tabletop from './resources/Pine_wood_Table_Top.jpg';
import TrickView from './TrickView';
import OwnCardsView from './OwnCardsView';
import PositionedPlayerView from './PositionedPlayerView';
import { MatchPhase } from 'api/karlchen_pb';

interface Props {
  match: Match | null;
  players: PlayingUsers;
}
export default function GameView({ match, players }: Props): React.ReactElement {
  const dRef = useRef<HTMLDivElement>(null);
  const [size, setSize] = useState([0, 0]);
  useEffect(() => {
    if (dRef.current) {
      const div = dRef.current;
      const updateSize = () => {
        if (!dRef?.current) {
          return;
        }
        setSize([div.clientWidth, div.clientHeight]);
      };
      updateSize();
      window.addEventListener('resize', updateSize);
      return () => window.removeEventListener('resize', updateSize);
    }
  }, [dRef]);
  if (!match) {
    return <h2>no match</h2>;
  }
  if (!match.game) {
    return <h2>no game</h2>;
  }

  const myTurn = match.turn === Pos.bottom;
  const inGame = match.phase === MatchPhase.GAME && match.game;
  return (
    <div
      style={{
        position: 'relative',
        width: '100%',
        height: '100%',
        backgroundImage: `url(${tabletop})`,
        backgroundSize: 'cover',
      }}
      ref={dRef}
    >
      {/*<DeclarationDialogContainer />*/}
      <div
        style={{
          position: 'absolute',
          bottom: 0,
          width: '100%',
          height: '100%',
          overflow: 'hidden',
        }}
      >
        {inGame && (
          <TrickView
            trick={match.game.currentTrick}
            cardWidth={size[0] / 6}
            center={['50%', '50%']}
          />
        )}
        <OwnCardsView
          cards={match.cards}
          cardWidth={size[0] / 5}
          onClick={() => undefined /*todo*/}
        />
        {[Pos.left, Pos.right, Pos.top, Pos.bottom].map((p) => (
          <PositionedPlayerView key={p} user={players[p]} pos={p} match={match} />
        ))}
      </div>
    </div>
  );
}
