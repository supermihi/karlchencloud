import { connect } from 'react-redux';
import LoginView from './LoginView';
import { createSelector } from '@reduxjs/toolkit';
import { selectAuth } from 'app/auth/slice';
import { forgetLogin } from 'app/auth/thunks';
import * as session from 'app/session';
import { MyUserData } from 'app/auth';

const mapStateToProps = createSelector(selectAuth, session.selectSession, (auth, session) => ({
  loading: Boolean(session.starting),
  error: session.error,
  currentLogin: auth.storedLogin as MyUserData,
}));
const mapDispatchToProps = {
  forgetLogin,
  login: session.startSession,
  resetError: session.actions.resetError,
};
export default connect(mapStateToProps, mapDispatchToProps)(LoginView);
