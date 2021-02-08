import { mapValues } from 'lodash';
import { User } from './core';

export enum Pos {
  bottom,
  left,
  top,
  right,
}

export type PlayerMap<T> = Record<Pos, T>;
export type PlayerIds = PlayerMap<string>;
export type PartialPlayerMap<T> = Partial<PlayerMap<T>>;

export function newPlayerMap<T>(init: (p: Pos) => T): PlayerMap<T> {
  return {
    [Pos.bottom]: init(Pos.bottom),
    [Pos.left]: init(Pos.left),
    [Pos.top]: init(Pos.top),
    [Pos.right]: init(Pos.right),
  };
}

export function nthNext(pos: Pos, n: number): Pos {
  return (pos + n) % 4;
}

export function nextPos(pos: Pos): Pos {
  return nthNext(pos, 1);
}

export function getPosition(players: PlayerIds, id: string): Pos {
  for (const pos of [Pos.bottom, Pos.left, Pos.right, Pos.top]) {
    if (players[pos] === id) {
      return pos;
    }
  }
  throw new Error(`no position for player with id ${id}`);
}

export type PlayingUsers = PlayerMap<User>;

export function toPlayerMap(players: PlayerIds, users: Record<string, User>): PlayingUsers {
  return mapValues(players, (id) => users[id]);
}
