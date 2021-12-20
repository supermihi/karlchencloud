import React from 'react';
import { Match } from 'model/match';
import tabletop from './resources/Pine_wood_Table_Top.jpg';
import { useResizeAwareRef } from 'shared/resizeHook';
import { Table } from 'model/table';
import MatchView from 'play/view/MatchView';
import InGameView from 'play/view/InGameView';
import { createSelector } from '@reduxjs/toolkit';
import { toUserMap } from 'model/core';
import { PlayerIds, toPlayerMap } from 'model/players';
import { AppDispatch } from 'state';

interface Props {
  match: Match | null;
  table: Table;
  dispatch: AppDispatch;
}

const selectUserMap = createSelector(
  (p: Props) => p.table.members,
  (m) => toUserMap(m)
);
const selectPlayerMap = createSelector(
  selectUserMap,
  (p) => p.match?.players ?? ({} as PlayerIds),
  (users, players) => toPlayerMap(players, users)
);
export default function PlayView(props: Props): React.ReactElement {
  const [ref, size] = useResizeAwareRef<HTMLDivElement>();
  const { table, match, dispatch } = props;

  return (
    <div
      style={{
        position: 'relative',
        width: '100%',
        height: '100%',
        backgroundImage: `url(${tabletop})`,
        backgroundSize: 'cover',
      }}
      ref={ref}
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
        {match?.game && <InGameView game={match.game} tableSize={size} />}
        {match && (
          <MatchView
            match={match}
            dispatch={dispatch}
            table={table}
            tableSize={size}
            players={selectPlayerMap(props)}
          />
        )}
      </div>
    </div>
  );
}
