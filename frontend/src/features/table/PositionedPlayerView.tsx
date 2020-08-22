import { User } from 'model/core';
import { isAuction, Match } from 'model/match';
import { Pos } from 'model/players';
import React from 'react';
import PlayerView, { Props as PlayerViewProps } from './PlayerView';
interface Props {
  user: User;
  pos: Pos;
  match: Match;
}
export default function PositionedPlayerView({ user, pos, match }: Props) {
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
  const bids = isAuction(match.details) ? [] : match.details.bids[pos] || [];
  const soloGame = isAuction(match.details)
    ? undefined
    : match.details.mode.soloist === user.id
    ? match.details.mode.type
    : undefined;
  return {
    user,
    turn,
    bids,
    soloGame,
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
