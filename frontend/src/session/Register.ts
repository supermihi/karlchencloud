import { connect } from 'react-redux';
import RegisterView from './RegisterView';
import { createSelector } from '@reduxjs/toolkit';
import { selectSession } from './selectors';
import { register } from './thunks/register';

const mapState = createSelector(selectSession, (session) => ({
  loading: session.loading,
  error: session.error,
}));

export default connect(mapState, { register })(RegisterView);
