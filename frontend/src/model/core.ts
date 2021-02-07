import * as api from "api/karlchen_pb";
import fromPairs from "lodash.frompairs";

export interface User {
  name: string;
  id: string;
  online?: boolean;
}

export function toUserMap(users: User[]): Record<string, User> {
  return fromPairs(users.map((u) => [u.id, u]));
}

export interface Card {
  suit: api.Suit;
  rank: api.Rank;
}
