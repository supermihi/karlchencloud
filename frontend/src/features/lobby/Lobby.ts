import { connect } from 'react-redux';
import LobbyView from './LobbyView';
import { createSelector } from '@reduxjs/toolkit';

import { selectLobby } from './slice';
import { selectGame } from 'app/game/slice';
import { createTable } from 'app/game/thunks';
import { selectSession } from 'app/session';
import { User } from 'model/core';

const mapState = createSelector(selectGame, selectLobby, selectSession, (s, l, sess) => ({
  activeTable: s.currentTable,
  suppliedInviteCode: l.suppliedInviteCode,
  me: sess.session as User,
}));
const mapDispatch = {
  createTable: createTable.thunk,
};
export default connect(mapState, mapDispatch)(LobbyView);
