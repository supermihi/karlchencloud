import { connect } from 'react-redux';
import LobbyView from './LobbyView';
import { createSelector } from '@reduxjs/toolkit';
import { selectSession } from '../session/selectors';
import { selectTable } from '../play/selectors';
import { User } from '../model/core';
import { createTable, startTable } from '../play/thunks';

const mapStateToProps = createSelector(selectTable, selectSession, (table, session) => ({
  activeTable: table,
  me: session.activeSession as User,
}));
const mapDispatchToProps = {
  createTable,
  startTable,
};
export default connect(mapStateToProps, mapDispatchToProps)(LobbyView);
