import React, { HTMLProps } from 'react';
import { Match, isAuction } from 'model/match';
import tabletop from './Pine_wood_Table_Top.jpg';
import OwnCardsView from './OwnCardsView';
import TrickView from './TrickView';

import { User } from 'model/core';
import { Pos } from 'model/players';
import { mapValues } from 'lodash';
import PositionedPlayerView from './PositionedPlayerView';

interface Props extends HTMLProps<HTMLDivElement> {
  match: Match;
  users: Record<string, User>;
}

export default function MatchView({ match, users, style, ...props }: Props) {
  const playerUsers: Record<Pos, User> = mapValues(match.players, (id) => users[id]);
  return (
    <div
      {...props}
      style={{
        ...style,
        position: 'relative',
        backgroundImage: `url(${tabletop})`,
        backgroundSize: 'cover',
      }}
    >
      <div
        style={{
          position: 'absolute',
          bottom: 0,
          width: '100%',
          height: '100%',
          overflow: 'hidden',
        }}
      >
        <OwnCardsView cards={match.cards} cardWidth={150} />
        {isAuction(match.details) ? null : (
          <TrickView trick={match.details.currentTrick} cardWidth={120} center={['50%', '50%']} />
        )}
        {[Pos.left, Pos.right, Pos.top, Pos.bottom].map((p) => (
          <PositionedPlayerView key={p} user={playerUsers[p]} pos={p} match={match} />
        ))}
      </div>
    </div>
  );
}
