import * as React from 'react';
import { TableState } from 'model/table';
import { toUserMap } from 'model/core';
import MatchView from './MatchView';

interface Props {
  table: TableState;
}

export default function TableView({ table: { match, table } }: Props) {
  if (!match) {
    return <span>'no match'</span>;
  }
  const users = toUserMap(table.players);
  return <MatchView match={match} users={users} />;
}
