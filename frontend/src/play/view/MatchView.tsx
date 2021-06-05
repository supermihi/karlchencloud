import React from 'react';
import { Match } from 'model/match';
import OwnCardsView from 'play/view/OwnCardsView';
import { Pos, toPlayerMap } from 'model/players';
import PositionedPlayerView from 'play/view/PositionedPlayerView';
import { Size } from 'shared/resizeHook';
import { Table } from 'model/table';
import { toUserMap } from 'model/core';

interface Props {
  match: Match;
  table: Table;
  tableSize: Size;
}

const MatchView: React.FC<Props> = ({ match, table, tableSize }) => {
  const players = toPlayerMap(match.players, toUserMap(table.members));
  return (
    <>
      <OwnCardsView
        cards={match.cards}
        cardWidth={tableSize.width / 5}
        onClick={() => undefined /*todo*/}
      />
      {[Pos.left, Pos.right, Pos.top, Pos.bottom].map((p) => (
        <PositionedPlayerView key={p} user={players[p]} pos={p} match={match} />
      ))}
    </>
  );
};
export default MatchView;
