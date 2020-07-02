import { connect } from 'react-redux';
import RegisterView from './RegisterView';
import { createSelector } from '@reduxjs/toolkit';
import { selectAuth } from 'app/auth/slice';
import { register } from 'app/auth/thunks';

const mapState = createSelector(selectAuth, ({ registering, registerError }) => ({
  loading: registering,
  error: registerError,
}));

export default connect(mapState, { register })(RegisterView);
