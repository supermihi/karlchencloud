import { Card } from 'model/core';
import * as api from 'api/karlchen_pb';

export function fromCard(c: Card): api.Card {
  const ans = new api.Card();
  ans.setSuit(c.suit);
  ans.setRank(c.rank);
  return ans;
}

export function newPlayCardRequest(c: Card, tableId: string): api.PlayCardRequest {
  const req = new api.PlayCardRequest();
  const pbCard = fromCard(c);
  req.setCard(pbCard);
  req.setTable(tableId);
  return req;
}
