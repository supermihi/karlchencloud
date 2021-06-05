import { connect } from 'react-redux';
import LoginView from './LoginView';
import { createSelector } from '@reduxjs/toolkit';
import { selectSession } from '../selectors';
import { login } from '../thunks/authenticate';
import { actions } from '../slice';

const mapStateToProps = createSelector(selectSession, (session) => ({
  loading: Boolean(session.loading),
  error: session.error,
}));
const mapDispatchToProps = {
  login,
  resetError: actions.resetError,
};
export default connect(mapStateToProps, mapDispatchToProps)(LoginView);
