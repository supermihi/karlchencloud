import { connect } from 'react-redux';
import LobbyView from './LobbyView';
import { createSelector } from '@reduxjs/toolkit';
import { selectSession } from 'app/session';
import { createTable } from 'app/session/lobby';

const mapState = createSelector(selectSession, (s) => ({
  activeTable: s.currentTable,
}));
const mapDispatch = {
  createTable,
};
export default connect(mapState, mapDispatch)(LobbyView);
