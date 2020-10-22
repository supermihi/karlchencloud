import TableView from './MainView';
import { connect } from 'react-redux';
import { createSelector } from '@reduxjs/toolkit';
import { playCard } from 'app/game/match';
import { selectCurrentTableOrThrow, selectMatch } from 'app/game/selectors';

const mapState = createSelector(selectCurrentTableOrThrow, selectMatch, (t, m) => ({
  table: t,
  match: m,
}));
const mapDispatch = {
  playCard,
};

export default connect(mapState, mapDispatch)(TableView);
