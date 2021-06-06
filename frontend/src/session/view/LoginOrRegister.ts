import { connect } from 'react-redux';
import LoginOrRegisterView from './LoginOrRegisterView';
import { createSelector } from '@reduxjs/toolkit';
import { selectSession } from '../selectors';
import { login, register } from '../thunks/authenticate';
import { actions } from '../slice';
import { SessionPhase } from '../model';

const mapStateToProps = createSelector(selectSession, (session) => ({
  loading: session.phase === SessionPhase.Starting,
  error: session.error,
}));
const mapDispatchToProps = {
  login,
  register,
  resetError: actions.resetError,
};
export default connect(mapStateToProps, mapDispatchToProps)(LoginOrRegisterView);
