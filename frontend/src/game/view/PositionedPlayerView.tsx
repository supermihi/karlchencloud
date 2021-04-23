import React from 'react';
import { User } from 'model/core';
import { Match } from 'model/match';
import { Pos } from 'model/players';
import PlayerView, { Props as PlayerViewProps } from './PlayerView';
interface Props {
  user: User;
  pos: Pos;
  match: Match;
}
export default function PositionedPlayerView({ user, pos, match }: Props): React.ReactElement {
  const plProps = getPlayerProps(match, pos, user);
  const posProps = getPositionProps(pos);
  return (
    <div style={{ position: 'absolute', ...posProps }}>
      <PlayerView {...plProps} />
    </div>
  );
}

function getPlayerProps(match: Match, pos: Pos, user: User): PlayerViewProps {
  const turn = match.turn === pos;
  const bids = match.game?.bids[pos] ?? [];
  const soloGame = !!match.game
      ? match.game.mode.soloist === pos
      ? match.game.mode.type
      : undefined
      : undefined;
  const declaration = match.auction?.declarations[pos];
  return {
    user,
    turn,
    bids,
    soloGame,
    declaration,
  };
}

function getPositionProps(pos: Pos) {
  switch (pos) {
    case Pos.left:
      return {
        left: '10px',
        top: '50%',
        transform: 'translate(0, -50%)',
      };
    case Pos.top:
      return {
        left: '50%',
        top: '10px',
        transform: 'translate(-50%, 0)',
      };
    case Pos.right:
      return {
        right: '10px',
        top: '50%',
        transform: 'translate(0, -50%)',
      };
    case Pos.bottom:
      return {
        right: '50%',
        bottom: '10px',
        transform: 'translate(-50%, 0)',
      };
  }
}