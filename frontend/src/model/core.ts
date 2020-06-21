import * as api from "api/karlchen_pb";
export interface User {
  name: string;
  id: string;
  online?: boolean;
}

export interface Card {
  suit: api.Suit;
  rank: api.Rank;
}
