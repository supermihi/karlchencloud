import React, { useCallback } from "react";
import { useDispatch, useSelector } from "react-redux";
import Signup from "./Signup";
import Login from "./Login";
import { selectLogin, register, login } from "core/login";

export default () => {
  const { error, loading, me, secret, loggedIn } = useSelector(selectLogin);
  const dispatch = useDispatch();
  const loginCb = useCallback(() => dispatch(login()), [dispatch]);
  const registerCb = useCallback((name) => dispatch(register(name)), [
    dispatch,
  ]);
  if (secret && !loggedIn) {
    return (
      <Login error={error} loading={loading} name={me.name} login={loginCb} />
    );
  }
  return <Signup error={error} loading={loading} signup={registerCb} />;
};
