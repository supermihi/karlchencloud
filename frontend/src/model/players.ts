export type Players = Record<Pos, string>;

export enum Pos {
  bottom,
  left,
  top,
  right,
}
export function nthNext(pos: Pos, n: number): Pos {
  return (pos + n) % 4;
}

export function getPosition(players: Players, id: string): Pos {
  for (const pos of [Pos.bottom, Pos.left, Pos.right, Pos.top]) {
    if (players[pos] === id) {
      return pos;
    }
  }
  throw new Error(`no position for player with id ${id}`);
}
