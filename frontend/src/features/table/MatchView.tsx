import React, { HTMLProps } from 'react';
import { Match, isAuction } from 'model/match';
import tabletop from './Pine_wood_Table_Top.jpg';
import OwnCardsView from './OwnCardsView';
import ImgTrickView from './TrickView';
import PlayerView from './PlayerView';
import { User } from 'model/core';

interface Props extends HTMLProps<HTMLDivElement> {
  match: Match;
  users: Record<string, User>;
}

export default function MatchView({ match, users, style, ...props }: Props) {
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
          <ImgTrickView
            trick={match.details.currentTrick}
            players={match.players}
            cardWidth={150}
            center={['50%', '50%']}
          />
        )}
        <PositionedPlayerView
          left="10px"
          top="50%"
          transform="translate(0, -50%)"
          user={users[match.players[1]]}
        />
        <PositionedPlayerView
          left="50%"
          top="10px"
          transform="translate(-50%, 0)"
          user={users[match.players[2]]}
        />
        <PositionedPlayerView
          right="10px"
          top="50%"
          transform="translate(0, -50%)"
          user={users[match.players[3]]}
        />
      </div>
    </div>
  );
}
interface PositionedPlayerViewProps {
  left?: string;
  top?: string;
  right?: string;
  transform?: string;
  user: User;
}
function PositionedPlayerView({ user, ...css }: PositionedPlayerViewProps) {
  return (
    <div style={{ position: 'absolute', ...css }}>
      <PlayerView user={user} />
    </div>
  );
}
