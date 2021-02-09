import { me } from 'mocks/players';
import { myTable } from 'mocks/table';
import React from 'react';
import LobbyView from './LobbyView';

export default (
  <LobbyView activeTable={myTable} me={me} createTable={() => null} startTable={() => null} />
);
