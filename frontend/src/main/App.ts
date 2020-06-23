import { selectLocation } from 'app/routing';
import { connect } from 'react-redux';
import AppView from './AppView';
import { createSelector } from '@reduxjs/toolkit';

const mapState = createSelector(selectLocation, (location) => ({ location }));
export default connect(mapState)(AppView);
