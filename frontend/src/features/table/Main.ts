import TableView from './MainView';
import { connect } from 'react-redux';
import { selectGame } from 'app/game/slice';
import { createSelector } from '@reduxjs/toolkit';

const mapState = createSelector(selectGame, (g) => ({
  table: g.currentTable as any,
}));

export default connect(mapState)(TableView);
