import { connect } from "react-redux";
import LoginView from "./LoginView";
import { createSelector } from "@reduxjs/toolkit";
import { selectAuth, forgetLogin } from "app/auth/slice";
import { selectSession, startSession } from "app/session";
import { LoginData } from "app/auth";

const mapStateToProps = createSelector(
  selectAuth,
  selectSession,
  (auth, session) => ({
    loading: session.starting,
    error: session.error,
    currentLogin: auth.storedLogin as LoginData,
  })
);
const mapDispatchToProps = {
  forgetLogin,
  login: startSession,
};
export default connect(mapStateToProps, mapDispatchToProps)(LoginView);
