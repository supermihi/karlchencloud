import { connect } from 'react-redux';
import LoginView from './LoginView';
import { createSelector } from '@reduxjs/toolkit';
import { selectAuth, forgetLogin } from 'app/auth/slice';
import * as session from 'app/session';
import { LoginData } from 'app/auth';

const mapStateToProps = createSelector(
  selectAuth,
  session.selectSession,
  (auth, session) => ({
    loading: session.starting,
    error: session.error,
    currentLogin: auth.storedLogin as LoginData,
  })
);
const mapDispatchToProps = {
  forgetLogin,
  login: session.startSession,
  resetError: session.actions.resetError,
};
export default connect(mapStateToProps, mapDispatchToProps)(LoginView);
