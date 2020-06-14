import React from "react";
import { connect } from "react-redux";
import Signup from "./Signup";
import Login from "./Login";
import {
  selectAuth,
  register,
  AuthState,
  tryLogin,
  forgetLogin,
} from "./slice";
import { LoginData } from "./api";
interface Props extends AuthState {
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
