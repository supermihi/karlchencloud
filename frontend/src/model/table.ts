import { User } from "./core";
import { MatchPhase, TableMember, TableData } from "api/karlchen_pb";

export interface Table {
  id: string;
  owner: string;
  invite?: string;
  players: User[];
  match?: Match;
}

export interface Match {
  Phase: MatchPhase;
}

export function toUser(member: TableMember): User {
  return { id: member.getUserId(), name: member.getName() };
}

export function toTable(t: TableData): Table {
  return {
    owner: t.getOwner(),
    id: t.getTableId(),
    players: t.getMembersList().map(toUser),
  };
}
