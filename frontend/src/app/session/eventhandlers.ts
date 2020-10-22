import * as api from 'api/karlchen_pb';
import { selectPlayers } from 'app/game/selectors';
import { AppThunk, AppDispatch, RootState } from 'app/store';
import { getCurrentTableState, toMatch, toMode } from 'model/apiconv';
import { Declaration, DeclareResult } from 'model/auction';
import { getPosition } from 'model/players';
import * as events from './events';
const { EventCase } = api.Event;

export const onEvent = (event: api.Event): AppThunk => (dispatch, getState) => {
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
    case EventCase.DECLARED:
      onDeclared(event.getDeclared() as api.Declaration, dispatch, getState);
      return;
    default:
      console.log(`unimplemented event: ${event}`);
  }
};

function onWelcome(userState: api.UserState, dispatch: AppDispatch) {
  const name = userState.getName();
  dispatch(events.sessionStarted(name));
  const table = getCurrentTableState(userState);
  dispatch(events.tableChanged(table));
}

function onMember(event: api.MemberEvent, dispatch: AppDispatch) {
  const name = event.getName();
  const id = event.getUserId();
  switch (event.getType()) {
    case api.MemberEventType.JOIN_TABLE:
      dispatch(events.memberJoined({ name, id }));
      return;
    case api.MemberEventType.LEAVE_TABLE:
      dispatch(events.memberLeft(id));
      return;
    case api.MemberEventType.GO_OFFLINE:
      dispatch(events.memberStatusChanged({ name, id, online: false }));
      return;
    case api.MemberEventType.GO_ONLINE:
      dispatch(events.memberStatusChanged({ name, id, online: true }));
      return;
  }
}

function onStart(ms: api.MatchState, dispatch: AppDispatch) {
  dispatch(events.matchStarted(toMatch(ms)));
}

function onDeclared(decl: api.Declaration, dispatch: AppDispatch, getState: () => RootState) {
  const apiMode = decl.getDefinedgamemode();
  const players = selectPlayers(getState());
  const mode = apiMode === undefined ? null : toMode(apiMode, players);
  const declaration: DeclareResult = {
    mode,
    player: getPosition(players, decl.getUserId()),
    declaration: decl.getVorbehalt() ? Declaration.vorbehalt : Declaration.gesund,
  };
  dispatch(events.playerDeclared(declaration));
}
