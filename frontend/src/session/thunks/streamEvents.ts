import * as api from 'api/karlchen_pb';
import { TablePhase } from 'api/karlchen_pb';
import { selectPlayers } from 'play/selectors';
import { Dispatchable } from 'state';
import {
  getCurrentTableState,
  toMatch,
  toTable,
  toPlayedCard,
  toDeclareResult,
} from 'model/apiconv';
import * as events from '../events';

const { EventCase } = api.Event;

export const createEventAction = (event: api.Event): Dispatchable => {
  switch (event.getEventCase()) {
    case EventCase.WELCOME:
      return createWelcomeAction(event.getWelcome() as api.UserState);
    case EventCase.MEMBER:
      return createMemberAction(event.getMember() as api.MemberEvent);
    case EventCase.START:
      return createStartAction(event.getStart() as api.MatchState);
    case EventCase.DECLARED:
      return createDeclaredAction(event.getDeclared() as api.Declaration);
    case EventCase.PLAYED_CARD:
      return createPlayedCardAction(event.getPlayedCard() as api.PlayedCard);
    case EventCase.NEW_TABLE:
      return createNewTableAction(event.getNewTable() as api.TableData);
    default:
      return events.notImplementedEvent(event.getEventCase().toString());
  }
};

function createWelcomeAction(userState: api.UserState): Dispatchable {
  const table = getCurrentTableState(userState);
  return (dispatch) => {
    dispatch(events.sessionStarted());
    dispatch(events.tableChanged(table));
  };
}

function createMemberAction(event: api.MemberEvent): Dispatchable {
  const name = event.getName();
  const id = event.getUserId();
  switch (event.getType()) {
    case api.MemberEventType.JOIN_TABLE:
      return events.memberJoined({ name, id });
    case api.MemberEventType.LEAVE_TABLE:
      return events.memberLeft(id);
    case api.MemberEventType.GO_OFFLINE:
      return events.memberStatusChanged({ name, id, online: false });
    case api.MemberEventType.GO_ONLINE:
      return events.memberStatusChanged({ name, id, online: true });
  }
}

function createStartAction(ms: api.MatchState) {
  return events.matchStarted(toMatch(ms));
}

function createDeclaredAction(decl: api.Declaration): Dispatchable {
  return (dispatch, getState) => {
    const players = selectPlayers(getState());
    dispatch(events.playerDeclared(toDeclareResult(decl, players)));
  };
}

function createPlayedCardAction(event: api.PlayedCard): Dispatchable {
  return (dispatch, getState) => {
    const players = selectPlayers(getState());
    const playedCard = toPlayedCard(event, players);
    dispatch(events.cardPlayed(playedCard));
  };
}

function createNewTableAction(t: api.TableData): Dispatchable {
  const table = toTable(t, TablePhase.NOT_STARTED);
  return (dispatch) => {
    dispatch(events.tableCreated(table));
  };
}
