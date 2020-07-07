import { connect } from 'react-redux';
import AcceptInviteDialog from './AcceptInviteDialogView';
import { clearInviteCode, selectLobby } from './slice';
import { joinTable } from 'app/game/thunks';

const mapDispatch = {
  clearInviteCode,
  joinTable,
};

export default connect(selectLobby, mapDispatch, (state, dispatch) => ({
  open: Boolean(state.suppliedInviteCode),
  confirm: () => dispatch.joinTable(state.suppliedInviteCode as string),
  reject: () => dispatch.clearInviteCode(),
}))(AcceptInviteDialog);
