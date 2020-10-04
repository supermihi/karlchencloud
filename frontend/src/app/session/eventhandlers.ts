import * as api from 'api/karlchen_pb';
import { AppThunk, AppDispatch } from 'app/store';
import { actions } from '.';
import { actions as gameActions } from '../game';
import { actions as tableActions } from '../game/table';
import { getCurrentTableState, toMatch } from 'model/apiconv';

const { EventCase } = api.Event;

export const onEvent = (event: api.Event): AppThunk => (dispatch) => {
  switch (event.getEventCase()) {
    case EventCase.WELCOME:
      onWelcome(event.getWelcome() as api.UserState, dispatch);
      return;
    case EventCase.MEMBER:
      onMember(event.getMember() as api.MemberEvent, dispatch);
      return;
    case EventCase.START:
      onStart(event.getStart() as api.MatchState, dispatch);
      return;
    default:
      console.log(`unimplemented event: ${event}`);
  }
};

function onWelcome(userState: api.UserState, dispatch: AppDispatch) {
  const name = userState.getName();
  dispatch(actions.sessionStarted(name));
  const table = getCurrentTableState(userState);
  dispatch(gameActions.currentTableChanged(table));
}

function onMember(event: api.MemberEvent, dispatch: AppDispatch) {
  const name = event.getName();
  const id = event.getUserId();
  switch (event.getType()) {
    case api.MemberEventType.JOIN_TABLE:
      dispatch(tableActions.memberJoined({ name, id }));
      return;
    case api.MemberEventType.LEAVE_TABLE:
      dispatch(tableActions.memberLeft(id));
      return;
    case api.MemberEventType.GO_OFFLINE:
      dispatch(tableActions.memberStatusChanged({ name, id, online: false }));
      return;
    case api.MemberEventType.GO_ONLINE:
      dispatch(tableActions.memberStatusChanged({ name, id, online: true }));
      return;
  }
}

function onStart(ms: api.MatchState, dispatch: AppDispatch) {
  dispatch(tableActions.matchStarted(toMatch(ms)));
}
