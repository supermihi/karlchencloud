import React from 'react';
import TableCard from './TableCard';
import * as players from 'mocks/players';
import * as table from 'mocks/table';
import { useSelect, useValue } from 'react-cosmos/fixture';

export default function Fixture(): React.ReactElement {
  const [own] = useValue('own table', { defaultValue: false });
  return (
    <TableCard
      me={players.me}
      table={own ? table.myTable : table.someTable}
      startTable={() => null}
    />
  );
}
