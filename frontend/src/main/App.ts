import { hot } from 'react-hot-loader/root';
import { selectLocation } from 'app/routing';
import { connect } from 'react-redux';
import AppView from './AppView';
import { createSelector } from '@reduxjs/toolkit';

const mapState = createSelector(selectLocation, (location) => ({ location }));
const App = connect(mapState)(AppView);
export default process.env.NODE_ENV === 'development' ? hot(App) : App;
