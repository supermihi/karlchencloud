import { TableId } from "./karlchen_pb";

export function tableId(id: string): TableId {
  const ans = new TableId();
  ans.setValue(id);
  return ans;
}
