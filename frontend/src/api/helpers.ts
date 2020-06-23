import { TableId, Timestamp } from './karlchen_pb';

export function tableId(id: string): TableId {
  const ans = new TableId();
  ans.setValue(id);
  return ans;
}

export function toDate(ts: Timestamp): Date {
  return new Date(ts.getNanos() / 1e6);
}
