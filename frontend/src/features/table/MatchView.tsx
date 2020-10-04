import React, { HTMLProps } from 'react';
import { Match, isAuction } from 'model/match';
import tabletop from './Pine_wood_Table_Top.jpg';
import OwnCardsView from './OwnCardsView';
import TrickView from './TrickView';
import { PlayerMap, Pos } from 'model/players';
import PositionedPlayerView from './PositionedPlayerView';
import { Card } from 'model/core';

interface Props extends HTMLProps<HTMLDivElement> {
  match: Match;
  players: PlayerMap;
  playCard: (card: Card) => void;
}

export default function MatchView({ match, players, style, playCard, ...props }: Props) {
  const myTurn = match.turn === Pos.bottom;
  return (
    <div
      {...props}
      style={{
        ...style,
        position: 'relative',
        width: '100%',
        height: '100%',
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
        {isAuction(match.details) ? null : (
          <TrickView trick={match.details.currentTrick} cardWidth={120} center={['50%', '50%']} />
        )}
        <OwnCardsView
          cards={match.cards}
          cardWidth={150}
          onClick={myTurn ? (card) => playCard(card) : undefined}
        />
        {[Pos.left, Pos.right, Pos.top, Pos.bottom].map((p) => (
          <PositionedPlayerView key={p} user={players[p]} pos={p} match={match} />
        ))}
      </div>
    </div>
  );
}
