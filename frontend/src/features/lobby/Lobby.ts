import { connect } from 'react-redux';
import LobbyView from './LobbyView';
import { createSelector } from '@reduxjs/toolkit';

import { selectLobby } from './slice';
import { selectGame } from 'app/game/slice';
import { createTable } from 'app/game/thunks';

const mapState = createSelector(selectGame, selectLobby, (s, l) => ({
  activeTable: s.currentTable,
  suppliedInviteCode: l.suppliedInviteCode,
}));
const mapDispatch = {
  createTable,
};
export default connect(mapState, mapDispatch)(LobbyView);
