import * as api from 'api/karlchen_pb';
import { selectPlayers } from 'game/selectors';
import { Dispatchable } from 'state';
import { getCurrentTableState, toCard, toMatch, toMode } from 'model/apiconv';
import { Declaration, DeclareResult } from 'model/auction';
import { getPosition } from 'model/players';
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
    default:
      return events.notImplementedEvent(event.getEventCase().toString());
  }
};

function createWelcomeAction(userState: api.UserState): Dispatchable {
  const name = userState.getName();
  const table = getCurrentTableState(userState);
  return (dispatch) => {
    dispatch(events.sessionStarted(name));
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
  const apiMode = decl.getDefinedgamemode();
  return (dispatch, getState) => {
    const players = selectPlayers(getState());
    const mode = apiMode === undefined ? null : toMode(apiMode, players);
    const declaration: DeclareResult = {
      mode,
      player: getPosition(players, decl.getUserId()),
      declaration: decl.getVorbehalt() ? Declaration.vorbehalt : Declaration.gesund,
    };
    dispatch(events.playerDeclared(declaration));
  };
}

function createPlayedCardAction(event: api.PlayedCard): Dispatchable {
  const card = toCard(event.getCard() as api.Card);
  const winner = event.hasTrickWinner() ? (event.getTrickWinner() as api.PlayerValue) : null;
  return (dispatch, getState) => {
    const players = selectPlayers(getState());
    const player = getPosition(players, event.getUserId());
    dispatch(
      events.cardPlayed({
        card,
        player,
        trickWinner: winner === null ? undefined : getPosition(players, winner.getUserId()),
      })
    );
  };
}
