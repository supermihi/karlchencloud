import { connect } from 'react-redux';
import LoginWithTokenView from './LoginWithTokenView';
import { createSelector } from '@reduxjs/toolkit';
import { selectSession } from '../selectors';
import { SessionPhase } from '../model';
import { forgetLoginThunk as forgetLogin } from 'session/thunks/authenticate';
import { startSession } from '../thunks/session';
import { actions } from '../slice';

const mapStateToProps = createSelector(selectSession, (session) => ({
  loading: session.phase === SessionPhase.Starting,
  error: session.error,
  name: session.userData?.name ?? '',
}));
const mapDispatchToProps = {
  forgetLogin,
  login: () => startSession(),
  resetError: actions.resetError,
};
export default connect(mapStateToProps, mapDispatchToProps)(LoginWithTokenView);
