import { connect } from 'react-redux';
import LoginView from './LoginView';
import { createSelector } from '@reduxjs/toolkit';
import { selectSession } from './selectors';
import { MyUserData } from './model';
import { forgetLogin } from 'session/thunks/register';
import { startSession } from './thunks/session';
import { actions } from './slice';

const mapStateToProps = createSelector(selectSession, (session) => ({
  loading: Boolean(session.loading),
  error: session.error,
  name: (session.storedLogin as MyUserData).name,
}));
const mapDispatchToProps = {
  forgetLogin,
  login: startSession,
  resetError: actions.resetError,
};
export default connect(mapStateToProps, mapDispatchToProps)(LoginView);
