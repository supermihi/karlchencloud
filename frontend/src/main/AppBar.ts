import { connect } from 'react-redux';
import { endSession, selectSession } from '../app/session';
import { createSelector } from '@reduxjs/toolkit';
import AppBarView from './AppBarView';
const mapState = createSelector(selectSession, (s) => ({
  loggedIn: s.session !== null,
}));
const mapDispatch = { logout: endSession };
export default connect(mapState, mapDispatch)(AppBarView);
