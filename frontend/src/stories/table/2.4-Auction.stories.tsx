import React from 'react';
import MatchView from 'features/table/MatchView';
import { Match, dummyGame } from 'model/match';
import { MatchPhase } from 'api/karlchen_pb';
import * as data from './mocks';
import { Pos, toPlayerMap } from 'model/players';
import { action } from '@storybook/addon-actions';
import DeclarationDialog from 'features/table/DeclarationDialog';
import { emptyAuction } from 'model/auction';

/* eslint import/no-anonymous-default-export: [2, {"allowObject": true}] */
export default {
  title: 'Match/Auction',
  component: DeclarationDialog,
};

const match: Match = {
  phase: MatchPhase.GAME,
  players: data.players,
  cards: data.fullHand,
  auction: emptyAuction(),
  game: dummyGame(),
  turn: Pos.bottom,
};
export const MatchTable = () => {
  return (
    <MatchView
      playCard={action('play card')}
      style={{ width: '60vw', height: '95vh' }}
      match={match}
      players={toPlayerMap(match.players, data.userMap)}
    />
  );
};
