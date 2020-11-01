import { connect } from 'react-redux';
import DeclarationDialog from './DeclarationDialog';
import { declare } from 'app/game/match';
import { selectCurrentMatchOrThrow } from 'app/game/selectors';
import { createSelector } from '@reduxjs/toolkit';
import { MatchPhase } from 'api/karlchen_pb';
import { Pos } from 'model/players';

const mapState = createSelector(selectCurrentMatchOrThrow, (m) => ({
  open: m.phase === MatchPhase.AUCTION && m.turn === Pos.bottom,
}));
const mapDispatch = {
  declare,
};
export default connect(mapState, mapDispatch)(DeclarationDialog);
