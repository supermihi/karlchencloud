import { TablePhase } from 'api/karlchen_pb';
import { Table } from 'model/table';

import * as players from './players';

export const someTable: Table = {
  id: 'table42',
  owner: players.left.id,
  members: players.users,
  created: 'gestern',
  phase: TablePhase.NOT_STARTED,
};

export const myTable: Table = { ...someTable, owner: players.me.id };
