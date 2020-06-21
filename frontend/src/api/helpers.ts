import { TableId, Empty } from "./karlchen_pb";

export function tableId(id: string): TableId {
  const ans = new TableId();
  ans.setValue(id);
  return ans;
}

export function newEmpty(): Empty {
  return new Empty();
}
