import { connect } from 'react-redux';
import PlayView from 'play/view/PlayView';
import { createSelector } from '@reduxjs/toolkit';
import { selectCurrentTableOrThrow, selectMatch } from 'play/selectors';

const mapStateToProps = createSelector(selectMatch, selectCurrentTableOrThrow, (match, table) => ({
  match,
  table,
}));
export default connect(mapStateToProps)(PlayView);
