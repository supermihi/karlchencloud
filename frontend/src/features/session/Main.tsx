import React from "react";
import { connect } from "react-redux";
import Signup from "./Signup";
import Login from "./Login";
import {
  selectAuth,
  register,
  SessionState,
  tryLogin,
  forgetLogin,
} from "../../core/session/slice";
import { LoginData } from "../../core/session/localstorage";
interface Props extends SessionState {
  tryLogin: (login: LoginData) => void;
  register: (name: string) => void;
  forgetLogin: () => void;
}

function AuthMain({
  error,
  loading,
  storedLogin,
  register,
  tryLogin,
  forgetLogin,
}: Props) {
  if (storedLogin !== null) {
    return (
      <Login
        currentLogin={storedLogin}
        login={tryLogin}
        {...{ error, loading, forgetLogin }}
      />
    );
  }
  return <Signup {...{ error, loading, register }} />;
}

export default connect(selectAuth, { tryLogin, register, forgetLogin })(
  AuthMain
);
