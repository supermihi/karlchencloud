import { connect } from "react-redux";
import RegisterView from "./RegisterView";
import { createSelector } from "@reduxjs/toolkit";
import { selectAuth, register } from "app/auth/slice";

const mapState = createSelector(
  selectAuth,
  ({ registering, registerError }) => ({
    loading: registering,
    error: registerError,
  })
);

export default connect(mapState, { register })(RegisterView);
