import { User, toUserMap } from 'model/core';
import { PlayerIds, Pos } from 'model/players';

export const me: User = { id: 'me', name: 'Ich', online: true };
export const left: User = { id: 'left', name: 'Spieler Links', online: false };
export const face: User = { id: 'face', name: 'Spieler Gegenüber', online: true };
export const right: User = { id: 'right', name: 'Spieler Rechts', online: true };

export const users = [right, me, face, left];
export const userMap = toUserMap(users);

export const players: PlayerIds = {
  [Pos.bottom]: me.id,
  [Pos.left]: left.id,
  [Pos.top]: face.id,
  [Pos.right]: right.id,
};
