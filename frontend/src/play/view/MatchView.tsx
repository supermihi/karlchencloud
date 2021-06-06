import React from 'react';
import { Match } from 'model/match';
import OwnCardsView from 'play/view/OwnCardsView';
import { PlayingUsers, Pos } from 'model/players';
import PositionedPlayerView from 'play/view/PositionedPlayerView';
import { Size } from 'shared/resizeHook';
import { Table } from 'model/table';
import DeclarationDialog from './DeclarationDialog';
import { AppDispatch } from 'state';
import { MatchPhase } from 'api/karlchen_pb';
import { playCard } from '../thunks';
import { Card } from 'model/core';

interface Props {
  match: Match;
  table: Table;
  tableSize: Size;
  players: PlayingUsers;
  dispatch: AppDispatch;
}

const MatchView: React.FC<Props> = ({ match, dispatch, tableSize, players }) => {
  const onClickCard =
    match.phase === MatchPhase.GAME ? (card: Card) => dispatch(playCard(card)) : undefined;
  return (
    <>
      <OwnCardsView cards={match.cards} cardWidth={tableSize.width / 5} onClick={onClickCard} />
      {[Pos.left, Pos.right, Pos.top, Pos.bottom].map((p) => (
        <PositionedPlayerView key={p} user={players[p]} pos={p} match={match} />
      ))}

      <DeclarationDialog match={match} dispatch={dispatch} />
    </>
  );
};

export default MatchView;
