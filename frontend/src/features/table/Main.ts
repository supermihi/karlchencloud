import TableView from './MainView';
import { connect } from 'react-redux';
import { selectGame } from 'app/game';
import { createSelector } from '@reduxjs/toolkit';
import { playCard } from 'app/game/match';

const mapState = createSelector(selectGame, (g) => ({
  table: g.currentTable as any,
}));
const mapDispatch = {
  playCard,
};

export default connect(mapState, mapDispatch)(TableView);
